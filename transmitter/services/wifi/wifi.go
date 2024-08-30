package wifi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/kilianp07/AthleteIQBox/transmitter/services"
	"github.com/kilianp07/AthleteIQBox/utils"
	logger "github.com/kilianp07/AthleteIQBox/utils/logger"
	wireless "github.com/kilianp07/go-wireless"
	"github.com/paypal/gatt"
)

type avAps struct {
	Aps []avAp `json:"aps"`
}

type avAp struct {
	Ssid   string `json:"ssid"`
	Signal int    `json:"signal"`
}

type wiFiConf struct {
	ServiceUUID     string `json:"service_uuid"`
	WlanUUID        string `json:"wlan_uuid"`
	WlanDescription string `json:"wlan_description"`
	TriggerScanUUID string `json:"trigger_scan_uuid"`
	TriggerScanDesc string `json:"trigger_scan_description"`
	AvApsUUID       string `json:"available_aps_uuid"`
	AvApsDesc       string `json:"available_aps_description"`
	Winterface      string `json:"wlan_interface"`
}

type WiFi struct {
	conf            wiFiConf
	service         *gatt.Service
	wlanChar        *gatt.Characteristic
	triggerScanChar *gatt.Characteristic
	avApsChar       *gatt.Characteristic
	trigger         chan bool
	wc              *wireless.Client
	aps             avAps
	logger          *logger.Logger
}

// NewWiFi initializes a new WiFi service.
func NewWiFi() services.Service {
	return &WiFi{
		conf:    wiFiConf{},
		trigger: make(chan bool, 1),
		logger:  logger.GetLogger("Wlan service"),
	}
}

func (w *WiFi) Conf() any {
	return &w.conf
}

// Configure sets up the gatt service and characteristics.
func (w *WiFi) Configure() (*gatt.Service, error) {
	var err error

	fmt.Println(w.conf)

	w.wc, err = wireless.NewClient(w.conf.Winterface)
	if err != nil {
		return nil, fmt.Errorf("failed to create wireless client: %v", err)
	}
	w.wc.ScanTimeout = time.Second * 3

	w.service = gatt.NewService(gatt.MustParseUUID(w.conf.ServiceUUID))

	// Wlan adder characteristic
	w.wlanChar = w.service.AddCharacteristic(gatt.MustParseUUID(w.conf.WlanUUID))
	w.wlanChar.HandleWriteFunc(w.handleWLANWrite)
	w.wlanChar.AddDescriptor(gatt.MustParseUUID("550576d2-6229-4f2f-870b-e15044aa6b9b")).SetValue(utils.StringToByte(w.conf.WlanDescription))

	// Trigger scan characteristic
	w.triggerScanChar = w.service.AddCharacteristic(gatt.MustParseUUID(w.conf.TriggerScanUUID))
	w.triggerScanChar.AddDescriptor(gatt.MustParseUUID("b44a091d-86fd-4ae8-ad9c-981053a6275e")).SetValue(utils.StringToByte(w.conf.TriggerScanDesc))
	w.triggerScanChar.HandleWriteFunc(w.handleWriteTrigger)

	// Available access points characteristic
	w.avApsChar = w.service.AddCharacteristic(gatt.MustParseUUID(w.conf.AvApsUUID))
	w.avApsChar.AddDescriptor(gatt.MustParseUUID("454a73e6-a660-4566-8eb0-111f8bd4ee98")).SetValue(utils.StringToByte(w.conf.AvApsDesc))
	w.avApsChar.HandleReadFunc(w.handleReadAps)

	// Log the UUIDs being added
	w.logger.Debugf("Service UUID: %s", w.conf.ServiceUUID)
	w.logger.Debugf("Wlan UUID: %s", w.conf.WlanUUID)
	w.logger.Debugf("Trigger Scan UUID: %s", w.conf.TriggerScanUUID)
	w.logger.Debugf("Available APs UUID: %s", w.conf.AvApsUUID)

	go w.Scan()
	return w.service, nil
}

// Update is a no-op for the WiFi service.
func (w *WiFi) Update(_ any) error {
	return nil
}

func (w *WiFi) handleWLANWrite(_ gatt.Request, incoming []byte) (status byte) {
	wifi := wiFi{}
	w.logger.Infof("Received WLAN write request:", string(incoming))
	if err := json.Unmarshal(incoming, &wifi); err != nil {
		w.logger.Errorf("Failed to unmarshal data:", err)
		return gatt.StatusUnexpectedError
	}
	w.logger.Debugf("Unmarshaled WiFi data: %+v", wifi)

	if err := w.addNetwork(wifi); err != nil {
		w.logger.Errorf("Failed to add network:", err)
		return gatt.StatusUnexpectedError
	}
	return gatt.StatusSuccess
}

func (w *WiFi) handleWriteTrigger(_ gatt.Request, incoming []byte) (status byte) {
	w.logger.Infof("Received trigger scan write request:", string(incoming))
	triggerValue := strings.ToLower(string(incoming)) == "true"
	w.trigger <- triggerValue // Send the trigger value
	return gatt.StatusSuccess
}

func (w *WiFi) handleReadAps(rsp gatt.ResponseWriter, rr *gatt.ReadRequest) {
	value, err := json.Marshal(w.aps)
	if err != nil {
		w.logger.Errorf("Failed to marshal APs:", err)
		rsp.SetStatus(gatt.StatusUnexpectedError)
		return
	}

	// Create a buffer and compact the JSON
	var buf bytes.Buffer
	if err := json.Compact(&buf, value); err != nil {
		w.logger.Errorf("Failed to compact JSON:", err)
		rsp.SetStatus(gatt.StatusUnexpectedError)
		return
	}

	// Split data into MTU-sized chunks if needed
	mtu := rr.Central.MTU() - 3 // Reserve some bytes for ATT overhead
	data := buf.Bytes()

	// Check if data fits in one MTU
	if len(data) <= mtu {
		if _, err := rsp.Write(data); err != nil {
			w.logger.Errorf("Failed to write APs response:", err)
			rsp.SetStatus(gatt.StatusUnexpectedError)
			return
		}
	} else {
		// Send data in chunks
		for i := 0; i < len(data); i += mtu {
			end := i + mtu
			if end > len(data) {
				end = len(data)
			}
			chunk := data[i:end]
			if _, err := rsp.Write(chunk); err != nil {
				w.logger.Errorf("Failed to write APs response:", err)
				rsp.SetStatus(gatt.StatusUnexpectedError)
				return
			}
		}
	}

	rsp.SetStatus(gatt.StatusSuccess)
}

// GetServiceUUID returns the service UUID.
func (w *WiFi) GetServiceUUID() string {
	return w.conf.ServiceUUID
}

// Scan continuously scans for access points when triggered.
func (w *WiFi) Scan() {
	for {
		if <-w.trigger {
			w.logger.Infof("Scanning for access points")
			aps, err := w.wc.Scan()
			if err != nil {
				w.logger.Errorf("Failed to scan:", err)
				continue
			}

			var data avAps
			for _, ap := range aps {
				data.Aps = append(data.Aps, avAp{
					Ssid:   ap.SSID,
					Signal: ap.Signal,
				})
			}
			w.logger.Infof("Found %d access points\n", len(data.Aps))
			w.aps = data
		}
	}
}

func (w *WiFi) addNetwork(d wiFi) error {
	var (
		net = wireless.NewNetwork(d.SSID, d.Password)
		err error
	)
	w.logger.Infof("Adding network:", d.SSID)

	_, err = w.wc.AddOrUpdateNetwork(net)
	if err != nil {
		return fmt.Errorf("failed to add network: %v", err)
	}

	_, err = w.wc.Connect(net)
	if err != nil {
		return fmt.Errorf("failed to connect to network: %v", err)
	}
	return nil
}
