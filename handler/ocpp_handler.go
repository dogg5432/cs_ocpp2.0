package handler

import (
	"github.com/dogg5432/cs_ocpp2.0/util"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/authorization"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/availability"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/data"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/diagnostics"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/display"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/firmware"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/meter"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/provisioning"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/reservation"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/security"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/transactions"
	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"

	"encoding/json"
	"fmt"
	"time"
)

type DataSample struct {
	SampleString string  `json:"sample_string"`
	SampleValue  float64 `json:"sample_value"`
}

const defaultHeartbeatInterval = 300

var logDefault = util.LogDefault

func (c *CSMSHandler) OnAuthorize(chargingStationID string, request *authorization.AuthorizeRequest) (response *authorization.AuthorizeResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("client with token %v authorized", request.IdToken)
	response = authorization.NewAuthorizationResponse(*types.NewIdTokenInfo(types.AuthorizationStatusAccepted))
	return
}

func (c *CSMSHandler) OnHeartbeat(chargingStationID string, request *availability.HeartbeatRequest) (response *availability.HeartbeatResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("heartbeat handled")
	response = availability.NewHeartbeatResponse(types.DateTime{Time: time.Now()})
	return
}

func (c *CSMSHandler) OnStatusNotification(chargingStationID string, request *availability.StatusNotificationRequest) (response *availability.StatusNotificationResponse, err error) {
	info, ok := c.ChargingStations[chargingStationID]
	if !ok {
		return nil, fmt.Errorf("unknown charging station %v", chargingStationID)
	}
	if request.ConnectorID > 0 {
		connectorInfo := info.getConnector(request.ConnectorID)
		connectorInfo.status = request.ConnectorStatus
		logDefault(chargingStationID, request.GetFeatureName()).Infof("connector %v updated status to %v", request.ConnectorID, request.ConnectorStatus)
	} else {
		logDefault(chargingStationID, request.GetFeatureName()).Infof("couldn't update status for invalid connector %v", request.ConnectorID)
	}
	response = availability.NewStatusNotificationResponse()
	return
}

func (c *CSMSHandler) OnDataTransfer(chargingStationID string, request *data.DataTransferRequest) (response *data.DataTransferResponse, err error) {
	var dataSample DataSample
	err = json.Unmarshal(request.Data.([]byte), &dataSample)
	if err != nil {
		logDefault(chargingStationID, request.GetFeatureName()).
			Errorf("invalid data received: %v", request.Data)
		return nil, err
	}
	logDefault(chargingStationID, request.GetFeatureName()).
		Infof("data received: %v, %v", dataSample.SampleString, dataSample.SampleValue)
	return data.NewDataTransferResponse(data.DataTransferStatusAccepted), nil
}

func (c *CSMSHandler) OnLogStatusNotification(chargingStationID string, request *diagnostics.LogStatusNotificationRequest) (response *diagnostics.LogStatusNotificationResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("log upload status: %v", request.Status)
	response = diagnostics.NewLogStatusNotificationResponse()
	return
}

func (c *CSMSHandler) OnNotifyCustomerInformation(chargingStationID string, request *diagnostics.NotifyCustomerInformationRequest) (response *diagnostics.NotifyCustomerInformationResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("data report for request %v: %v", request.RequestID, request.Data)
	response = diagnostics.NewNotifyCustomerInformationResponse()
	return
}

func (c *CSMSHandler) OnNotifyEvent(chargingStationID string, request *diagnostics.NotifyEventRequest) (response *diagnostics.NotifyEventResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("report part %v for events:\n", request.SeqNo)
	for _, ed := range request.EventData {
		logDefault(chargingStationID, request.GetFeatureName()).Infof("%v", ed)
	}
	response = diagnostics.NewNotifyEventResponse()
	return
}

func (c *CSMSHandler) OnNotifyMonitoringReport(chargingStationID string, request *diagnostics.NotifyMonitoringReportRequest) (response *diagnostics.NotifyMonitoringReportResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("report part %v for monitored variables:\n", request.SeqNo)
	for _, md := range request.Monitor {
		logDefault(chargingStationID, request.GetFeatureName()).Infof("%v", md)
	}
	response = diagnostics.NewNotifyMonitoringReportResponse()
	return
}

func (c *CSMSHandler) OnNotifyDisplayMessages(chargingStationID string, request *display.NotifyDisplayMessagesRequest) (response *display.NotifyDisplayMessagesResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("received display messages for request %v:\n", request.RequestID)
	for _, msg := range request.MessageInfo {
		logDefault(chargingStationID, request.GetFeatureName()).Printf("%v", msg)
	}
	response = display.NewNotifyDisplayMessagesResponse()
	return
}

func (c *CSMSHandler) OnFirmwareStatusNotification(chargingStationID string, request *firmware.FirmwareStatusNotificationRequest) (response *firmware.FirmwareStatusNotificationResponse, err error) {
	info, ok := c.ChargingStations[chargingStationID]
	if !ok {
		err = fmt.Errorf("unknown charging station %v", chargingStationID)
		return
	}
	info.FirmwareStatus = request.Status
	logDefault(chargingStationID, request.GetFeatureName()).Infof("updated firmware status to %v", request.Status)
	response = firmware.NewFirmwareStatusNotificationResponse()
	return
}

func (c *CSMSHandler) OnPublishFirmwareStatusNotification(chargingStationID string, request *firmware.PublishFirmwareStatusNotificationRequest) (response *firmware.PublishFirmwareStatusNotificationResponse, err error) {
	if len(request.Location) > 0 {
		logDefault(chargingStationID, request.GetFeatureName()).Infof("firmware download status on local controller: %v, download locations: %v", request.Status, request.Location)
	} else {
		logDefault(chargingStationID, request.GetFeatureName()).Infof("firmware download status on local controller: %v", request.Status)
	}
	response = firmware.NewPublishFirmwareStatusNotificationResponse()
	return
}

func (c *CSMSHandler) OnMeterValues(chargingStationID string, request *meter.MeterValuesRequest) (response *meter.MeterValuesResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("received meter values for EVSE %v. Meter values:\n", request.EvseID)
	for _, mv := range request.MeterValue {
		logDefault(chargingStationID, request.GetFeatureName()).Printf("%v", mv)
	}
	response = meter.NewMeterValuesResponse()
	return
}

func (c *CSMSHandler) OnBootNotification(chargingStationID string, request *provisioning.BootNotificationRequest) (response *provisioning.BootNotificationResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("boot confirmed for %v %v, serial: %v, firmare version: %v, reason: %v",
		request.ChargingStation.VendorName, request.ChargingStation.Model, request.ChargingStation.SerialNumber, request.ChargingStation.FirmwareVersion, request.Reason)
	response = provisioning.NewBootNotificationResponse(types.NewDateTime(time.Now()), defaultHeartbeatInterval, provisioning.RegistrationStatusAccepted)
	return
}

func (c *CSMSHandler) OnNotifyReport(chargingStationID string, request *provisioning.NotifyReportRequest) (response *provisioning.NotifyReportResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("data report %v, seq. %v:\n", request.RequestID, request.SeqNo)
	for _, d := range request.ReportData {
		logDefault(chargingStationID, request.GetFeatureName()).Printf("%v", d)
	}
	response = provisioning.NewNotifyReportResponse()
	return
}

func (c *CSMSHandler) OnReservationStatusUpdate(chargingStationID string, request *reservation.ReservationStatusUpdateRequest) (response *reservation.ReservationStatusUpdateResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("updated status of reservation %v to: %v", request.ReservationID, request.Status)
	response = reservation.NewReservationStatusUpdateResponse()
	return
}

func (c *CSMSHandler) OnSecurityEventNotification(chargingStationID string, request *security.SecurityEventNotificationRequest) (response *security.SecurityEventNotificationResponse, err error) {
	logDefault(chargingStationID, request.GetFeatureName()).Infof("type: %s, info: %s", request.Type, request.TechInfo)
	response = security.NewSecurityEventNotificationResponse()
	return
}

func (c *CSMSHandler) OnTransactionEvent(chargingStationID string, request *transactions.TransactionEventRequest) (response *transactions.TransactionEventResponse, err error) {
	switch request.EventType {
	case transactions.TransactionEventStarted:
		logDefault(chargingStationID, request.GetFeatureName()).Infof("transaction %v started, reason: %v, state: %v", request.TransactionInfo.TransactionID, request.TriggerReason, request.TransactionInfo.ChargingState)
	case transactions.TransactionEventUpdated:
		logDefault(chargingStationID, request.GetFeatureName()).Infof("transaction %v updated, reason: %v, state: %v\n", request.TransactionInfo.TransactionID, request.TriggerReason, request.TransactionInfo.ChargingState)
		for _, mv := range request.MeterValue {
			logDefault(chargingStationID, request.GetFeatureName()).Printf("%v", mv)
		}
	case transactions.TransactionEventEnded:
		logDefault(chargingStationID, request.GetFeatureName()).Infof("transaction %v stopped, reason: %v, state: %v\n", request.TransactionInfo.TransactionID, request.TriggerReason, request.TransactionInfo.ChargingState)
	}
	response = transactions.NewTransactionEventResponse()
	return
}
