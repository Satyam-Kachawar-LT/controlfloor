package main

import (
	"encoding/json"
	"fmt"

	uj "github.com/LambdaTest/ujsonin_private/v2/mod"
)

type ProvBase interface {
	asText(int16) string
	needsResponse() bool
	resHandler() func(uj.JNode, []byte)
}

type ProvPing struct {
	blah  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvPing) resHandler() func(uj.JNode, []byte) { return self.onRes }
func (self *ProvPing) needsResponse() bool                { return true }
func (self *ProvPing) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"ping\"}\n", id)
}

type ProvDoubleclick struct {
	udid  string
	x     int
	y     int
	onRes func(uj.JNode, []byte)
}

func (self *ProvDoubleclick) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvDoubleclick) needsResponse() bool { return true }
func (self *ProvDoubleclick) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"doubleclick\",udid:\"%s\",x:%d,y:%d}\n", id, self.udid, self.x, self.y)
}

type ProvRotate struct {
	udid       string
	isPortrait bool
	onRes      func(uj.JNode, []byte)
}

type ProvClick struct {
	udid  string
	x     int
	y     int
	onRes func(uj.JNode, []byte)
}

func (self *ProvRotate) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvRotate) needsResponse() bool { return true }
func (self *ProvRotate) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"rotate\",udid:\"%s\",isPortrait:%v}\n", id, self.udid, self.isPortrait)
}

type ProvAutoAcceptAlerts struct {
	udid             string
	autoAcceptAlerts string
	onRes            func(uj.JNode, []byte)
}

func (self *ProvAutoAcceptAlerts) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvAutoAcceptAlerts) needsResponse() bool { return true }
func (self *ProvAutoAcceptAlerts) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"autoAcceptAlerts\",udid:\"%s\",autoAcceptAlerts:\"%s\"}\n", id, self.udid, self.autoAcceptAlerts)
}

func (self *ProvClick) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvClick) needsResponse() bool { return true }
func (self *ProvClick) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"click\",udid:\"%s\",x:%d,y:%d}\n", id, self.udid, self.x, self.y)
}

type ProvLaunch struct {
	udid  string
	bid   string
	onRes func(uj.JNode, []byte)
}

func (self *ProvLaunch) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvLaunch) needsResponse() bool { return true }
func (self *ProvLaunch) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"launch\",udid:\"%s\",bid:\"%s\"}\n", id, self.udid, self.bid)
}

type ProvKill struct {
	udid  string
	bid   string
	onRes func(uj.JNode, []byte)
}

func (self *ProvKill) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvKill) needsResponse() bool { return true }
func (self *ProvKill) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"kill\",udid:\"%s\",bid:\"%s\"}\n", id, self.udid, self.bid)
}

type ProvAllowApp struct {
	udid  string
	bid   string
	onRes func(uj.JNode, []byte)
}

func (self *ProvAllowApp) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvAllowApp) needsResponse() bool { return true }
func (self *ProvAllowApp) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"allowApp\",udid:\"%s\",bid:\"%s\"}\n", id, self.udid, self.bid)
}

type ProvRestrictApp struct {
	udid  string
	bid   string
	onRes func(uj.JNode, []byte)
}

func (self *ProvRestrictApp) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvRestrictApp) needsResponse() bool { return true }
func (self *ProvRestrictApp) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"restrictApp\",udid:\"%s\",bid:\"%s\"}\n", id, self.udid, self.bid)
}

type ProvListRestrictedApps struct {
	udid  string
	bid   string
	onRes func(uj.JNode, []byte)
}

func (self *ProvListRestrictedApps) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvListRestrictedApps) needsResponse() bool { return true }
func (self *ProvListRestrictedApps) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"listRestrictedApps\",udid:\"%s\"}\n", id, self.udid, self.bid)
}

type ProvMouseDown struct {
	udid  string
	x     int
	y     int
	onRes func(uj.JNode, []byte)
}

func (self *ProvMouseDown) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvMouseDown) needsResponse() bool { return true }
func (self *ProvMouseDown) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"mouseDown\",udid:\"%s\",x:%d,y:%d}\n", id, self.udid, self.x, self.y)
}

type ProvMouseUp struct {
	udid  string
	x     int
	y     int
	onRes func(uj.JNode, []byte)
}

func (self *ProvMouseUp) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvMouseUp) needsResponse() bool { return true }
func (self *ProvMouseUp) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"mouseUp\",udid:\"%s\",x:%d,y:%d}\n", id, self.udid, self.x, self.y)
}

type ProvHardPress struct {
	udid string
	x    int
	y    int
}

func (self *ProvHardPress) resHandler() func(uj.JNode, []byte) { return nil }
func (self *ProvHardPress) needsResponse() bool                { return false }
func (self *ProvHardPress) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"hardPress\",udid:\"%s\",x:%d,y:%d}\n", id, self.udid, self.x, self.y)
}

type ProvInitWebrtcMsg struct {
	Id    int16  `json:"id"`
	Type  string `json:"type"`
	Udid  string `json:"udid"`
	Offer string `json:"offer"`
}
type ProvInitWebrtc struct {
	udid  string
	offer string
	onRes func(uj.JNode, []byte)
}

func (self *ProvInitWebrtc) resHandler() func(uj.JNode, []byte) {
	return self.onRes
}
func (self *ProvInitWebrtc) needsResponse() bool { return true }
func (self *ProvInitWebrtc) asText(id int16) string {
	msg := ProvInitWebrtcMsg{
		Id:    id,
		Type:  "initWebrtc",
		Udid:  self.udid,
		Offer: self.offer,
	}
	res, _ := json.Marshal(msg)
	return string(res)
}

type ProvLongPress struct {
	udid  string
	x     int
	y     int
	time  float64
	onRes func(uj.JNode, []byte)
}

func (self *ProvLongPress) resHandler() func(uj.JNode, []byte) {
	return self.onRes
}
func (self *ProvLongPress) needsResponse() bool { return true }
func (self *ProvLongPress) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"longPress\",udid:\"%s\",x:%d,y:%d,time:\"%f\"}\n", id, self.udid, self.x, self.y, self.time)
}

type ProvHome struct {
	udid  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvHome) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvHome) needsResponse() bool { return true }
func (self *ProvHome) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"home\",udid:\"%s\"}\n", id, self.udid)
}

type ProvShake struct {
	udid  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvShake) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvShake) needsResponse() bool { return true }
func (self *ProvShake) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"shake\",udid:\"%s\"}\n", id, self.udid)
}

type ProvCC struct {
	udid  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvCC) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvCC) needsResponse() bool { return true }
func (self *ProvCC) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"cc\",udid:\"%s\"}\n", id, self.udid)
}

type ProvAssistiveTouch struct {
	udid  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvAssistiveTouch) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvAssistiveTouch) needsResponse() bool { return true }
func (self *ProvAssistiveTouch) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"assistiveTouch\",udid:\"%s\"}\n", id, self.udid)
}

type ProvTaskSwitcher struct {
	udid  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvTaskSwitcher) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvTaskSwitcher) needsResponse() bool { return true }
func (self *ProvTaskSwitcher) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"taskSwitcher\",udid:\"%s\"}\n", id, self.udid)
}

type ProvWifiIp struct {
	udid  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvWifiIp) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvWifiIp) needsResponse() bool { return true }
func (self *ProvWifiIp) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"wifiIp\",udid:\"%s\"}\n", id, self.udid)
}

type ProvRefresh struct {
	udid  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvRefresh) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvRefresh) needsResponse() bool { return true }
func (self *ProvRefresh) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"refresh\",udid:\"%s\"}\n", id, self.udid)
}

type ProvRestart struct {
	udid  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvRestart) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvRestart) needsResponse() bool { return true }
func (self *ProvRestart) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"restart\",udid:\"%s\"}\n", id, self.udid)
}

type ProvSource struct {
	udid  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvSource) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvSource) needsResponse() bool { return true }
func (self *ProvSource) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"source\",udid:\"%s\"}\n", id, self.udid)
}

type ProvShutdown struct {
	onRes func(uj.JNode, []byte)
}

func (self *ProvShutdown) resHandler() func(uj.JNode, []byte) { return nil }
func (self *ProvShutdown) needsResponse() bool                { return false }
func (self *ProvShutdown) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"shutdown\"}\n", id)
}

type ProvKeys struct {
	udid     string
	keys     string
	curid    int
	prevkeys string
	onRes    func(uj.JNode, []byte)
}

func (self *ProvKeys) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvKeys) needsResponse() bool { return true }
func (self *ProvKeys) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"keys\",udid:\"%s\",keys:\"%s\",curid:%d,prevkeys:\"%s\"}\n",
		id, self.udid, self.keys, self.curid, self.prevkeys)
}

type ProvTestMsg struct {
	Id   int16  `json:"id"`
	Type string `json:"type"`
	Udid string `json:"udid"`
	Text string `json:"text"`
}
type ProvText struct {
	udid  string
	text  string
	onRes func(uj.JNode, []byte)
}

func (self *ProvText) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvText) needsResponse() bool { return true }
func (self *ProvText) asText(id int16) string {
	msg := ProvTestMsg{
		Id:   id,
		Type: "text",
		Udid: self.udid,
		Text: self.text,
	}
	res, _ := json.Marshal(msg)
	return string(res)
}

type ProvSwipe struct {
	udid  string
	x1    int
	y1    int
	x2    int
	y2    int
	delay float64
	onRes func(uj.JNode, []byte)
}

func (self *ProvSwipe) resHandler() func(data uj.JNode, rawData []byte) {
	return self.onRes
}
func (self *ProvSwipe) needsResponse() bool { return true }
func (self *ProvSwipe) asText(id int16) string {
	delayBy100 := int(self.delay * 100)
	return fmt.Sprintf("{id:%d,type:\"swipe\",udid:\"%s\",x1:%d,y1:%d,x2:%d,y2:%d,delay:%d}\n",
		id, self.udid, self.x1, self.y1, self.x2, self.y2, delayBy100)
}

type ProvStartStream struct {
	udid string
}

func (self *ProvStartStream) resHandler() func(uj.JNode, []byte) { return nil }
func (self *ProvStartStream) needsResponse() bool                { return false }
func (self *ProvStartStream) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"startStream\",udid:\"%s\"}\n", id, self.udid)
}

type ProvStopStream struct {
	udid string
}

func (self *ProvStopStream) resHandler() func(uj.JNode, []byte) {
	return nil
}

func (self *ProvStopStream) asText(id int16) string {
	return fmt.Sprintf("{id:%d,type:\"stopStream\",udid:\"%s\"}\n", id, self.udid)
}

func (self *ProvStopStream) needsResponse() bool {
	return false
}
