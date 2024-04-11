package answer

import (
	"github.com/bettercallmolly/belfast/connection"

	"github.com/bettercallmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

var validSC19001 protobuf.SC_19001

func DormData(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_19001 // Send an empty DormData
	response.Lv = proto.Uint32(0)
	response.Food = proto.Uint32(0)
	response.FoodMaxIncrease = proto.Uint32(0)
	response.FoodMaxIncreaseCount = proto.Uint32(0)
	response.FloorNum = proto.Uint32(0)
	response.ExpPos = proto.Uint32(0)
	response.NextTimestamp = proto.Uint32(0)
	response.LoadExp = proto.Uint32(0)
	response.LoadFood = proto.Uint32(0)
	response.LoadTime = proto.Uint32(0)
	response.Name = proto.String("")
	return client.SendMessage(19001, &response)
}

func init() {
	data := []byte{0x08, 0x04, 0x10, 0x00, 0x18, 0x00, 0x20, 0x00, 0x28, 0xc3, 0x1a, 0x28, 0x8c, 0x11, 0x28, 0xe7, 0x05, 0x28, 0x82, 0x0e, 0x32, 0x05, 0x08, 0xe9, 0x07, 0x10, 0x01, 0x32, 0x05, 0x08, 0xea, 0x07, 0x10, 0x01, 0x32, 0x05, 0x08, 0xce, 0x08, 0x10, 0x01, 0x32, 0x05, 0x08, 0xcf, 0x08, 0x10, 0x01, 0x32, 0x05, 0x08, 0xd0, 0x08, 0x10, 0x01, 0x32, 0x05, 0x08, 0xd1, 0x08, 0x10, 0x01, 0x32, 0x05, 0x08, 0xd2, 0x08, 0x10, 0x01, 0x32, 0x05, 0x08, 0xd3, 0x08, 0x10, 0x01, 0x32, 0x05, 0x08, 0xd7, 0x08, 0x10, 0x01, 0x32, 0x05, 0x08, 0xd8, 0x08, 0x10, 0x01, 0x32, 0x05, 0x08, 0x95, 0x0a, 0x10, 0x01, 0x32, 0x05, 0x08, 0x96, 0x0a, 0x10, 0x01, 0x32, 0x05, 0x08, 0x98, 0x0a, 0x10, 0x01, 0x32, 0x05, 0x08, 0x9a, 0x0a, 0x10, 0x01, 0x32, 0x05, 0x08, 0xba, 0x10, 0x10, 0x01, 0x32, 0x05, 0x08, 0x9e, 0x11, 0x10, 0x01, 0x32, 0x05, 0x08, 0x8a, 0x20, 0x10, 0x01, 0x32, 0x05, 0x08, 0xf9, 0x55, 0x10, 0x01, 0x32, 0x05, 0x08, 0xc6, 0x5e, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe0, 0x85, 0x01, 0x10, 0x01, 0x32, 0x06, 0x08, 0x82, 0x93, 0x02, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd7, 0xa9, 0x02, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa1, 0xb2, 0x02, 0x10, 0x01, 0x32, 0x06, 0x08, 0xf9, 0xcc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xfa, 0xcc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdd, 0xcd, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdf, 0xcd, 0x03, 0x10, 0x04, 0x32, 0x06, 0x08, 0xe0, 0xcd, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe1, 0xcd, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe2, 0xcd, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe3, 0xcd, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe4, 0xcd, 0x03, 0x10, 0x02, 0x32, 0x06, 0x08, 0xe5, 0xcd, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe7, 0xcd, 0x03, 0x10, 0x02, 0x32, 0x06, 0x08, 0xe8, 0xcd, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe9, 0xcd, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xea, 0xcd, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa5, 0xcf, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa6, 0xcf, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa7, 0xcf, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe9, 0xfb, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xea, 0xfb, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xcd, 0xfc, 0x03, 0x10, 0x04, 0x32, 0x06, 0x08, 0xce, 0xfc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xcf, 0xfc, 0x03, 0x10, 0x03, 0x32, 0x06, 0x08, 0xd0, 0xfc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd1, 0xfc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd2, 0xfc, 0x03, 0x10, 0x02, 0x32, 0x06, 0x08, 0xd3, 0xfc, 0x03, 0x10, 0x02, 0x32, 0x06, 0x08, 0xd4, 0xfc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd5, 0xfc, 0x03, 0x10, 0x03, 0x32, 0x06, 0x08, 0xd6, 0xfc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd7, 0xfc, 0x03, 0x10, 0x02, 0x32, 0x06, 0x08, 0xd8, 0xfc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd9, 0xfc, 0x03, 0x10, 0x02, 0x32, 0x06, 0x08, 0xda, 0xfc, 0x03, 0x10, 0x05, 0x32, 0x06, 0x08, 0xdb, 0xfc, 0x03, 0x10, 0x03, 0x32, 0x06, 0x08, 0xdc, 0xfc, 0x03, 0x10, 0x02, 0x32, 0x06, 0x08, 0xdd, 0xfc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xde, 0xfc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdf, 0xfc, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0x95, 0xfe, 0x03, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd0, 0xf9, 0x04, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd5, 0xf9, 0x04, 0x10, 0x01, 0x32, 0x06, 0x08, 0x95, 0xfb, 0x04, 0x10, 0x01, 0x32, 0x06, 0x08, 0x9a, 0xfb, 0x04, 0x10, 0x01, 0x32, 0x06, 0x08, 0xda, 0xa7, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xb5, 0xb0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa9, 0xb7, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xaa, 0xb7, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x8d, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x8e, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x8f, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x90, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x91, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x92, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x93, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x94, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x95, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x96, 0xb8, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0x98, 0xb8, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0x9a, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x9b, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x9c, 0xb8, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0x9e, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x9f, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa0, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa1, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa2, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa3, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa4, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xa5, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xf1, 0xb8, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd5, 0xb9, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0xd7, 0xb9, 0x05, 0x10, 0x04, 0x32, 0x06, 0x08, 0xdb, 0xb9, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdc, 0xb9, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdd, 0xb9, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x91, 0xbf, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x92, 0xbf, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xf5, 0xbf, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xf6, 0xbf, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xf7, 0xbf, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xf9, 0xbf, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xfa, 0xbf, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xfb, 0xbf, 0x05, 0x10, 0x04, 0x32, 0x06, 0x08, 0xff, 0xbf, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x80, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x81, 0xc0, 0x05, 0x10, 0x03, 0x32, 0x06, 0x08, 0x84, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x86, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x87, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x88, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x89, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x8a, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x8b, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x8c, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x8d, 0xc0, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0x8f, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x90, 0xc0, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe9, 0xf5, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xea, 0xf5, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xcd, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xce, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xcf, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd0, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd1, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd2, 0xf6, 0x05, 0x10, 0x04, 0x32, 0x06, 0x08, 0xd6, 0xf6, 0x05, 0x10, 0x06, 0x32, 0x06, 0x08, 0xdc, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdd, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xde, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdf, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe0, 0xf6, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0xe2, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe3, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe4, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe5, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe6, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe7, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe8, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xe9, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xea, 0xf6, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x95, 0xf8, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0xd1, 0xfd, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd2, 0xfd, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xb5, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xb6, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xb7, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xb8, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xb9, 0xfe, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0xbb, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xbc, 0xfe, 0x05, 0x10, 0x03, 0x32, 0x06, 0x08, 0xbf, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc0, 0xfe, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0xc2, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc3, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc4, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc5, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc6, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc7, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc8, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc9, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xca, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xcb, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xcc, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xcd, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xce, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xcf, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd0, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd1, 0xfe, 0x05, 0x10, 0x07, 0x32, 0x06, 0x08, 0xd8, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xd9, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xda, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdb, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdc, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xdd, 0xfe, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x99, 0xff, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0x9a, 0xff, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xfd, 0xff, 0x05, 0x10, 0x01, 0x32, 0x06, 0x08, 0xfe, 0xff, 0x05, 0x10, 0x02, 0x32, 0x06, 0x08, 0x80, 0x80, 0x06, 0x10, 0x08, 0x32, 0x06, 0x08, 0x88, 0x80, 0x06, 0x10, 0x01, 0x32, 0x06, 0x08, 0x89, 0x80, 0x06, 0x10, 0x01, 0x32, 0x06, 0x08, 0x8a, 0x80, 0x06, 0x10, 0x01, 0x32, 0x06, 0x08, 0xad, 0x8d, 0x06, 0x10, 0x01, 0x32, 0x06, 0x08, 0xbb, 0x8d, 0x06, 0x10, 0x01, 0x32, 0x06, 0x08, 0xbc, 0x8d, 0x06, 0x10, 0x01, 0x32, 0x06, 0x08, 0xbd, 0x8d, 0x06, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc2, 0x8d, 0x06, 0x10, 0x01, 0x32, 0x06, 0x08, 0xc3, 0x8d, 0x06, 0x10, 0x01, 0x32, 0x05, 0x08, 0xfb, 0x01, 0x10, 0x01, 0x32, 0x05, 0x08, 0xf2, 0x01, 0x10, 0x01, 0x32, 0x0b, 0x08, 0xe2, 0x01, 0x10, 0x01, 0x18, 0xb4, 0xfa, 0xa8, 0x95, 0x06, 0x32, 0x05, 0x08, 0xde, 0x01, 0x10, 0x01, 0x32, 0x0b, 0x08, 0xb7, 0x01, 0x10, 0x01, 0x18, 0xdd, 0xd6, 0xd6, 0x96, 0x06, 0x32, 0x04, 0x08, 0x6b, 0x10, 0x01, 0x32, 0x04, 0x08, 0x69, 0x10, 0x01, 0x32, 0x04, 0x08, 0x68, 0x10, 0x01, 0x32, 0x04, 0x08, 0x66, 0x10, 0x01, 0x32, 0x04, 0x08, 0x19, 0x10, 0x01, 0x32, 0x04, 0x08, 0x0e, 0x10, 0x01, 0x32, 0x04, 0x08, 0x0d, 0x10, 0x01, 0x32, 0x04, 0x08, 0x0c, 0x10, 0x01, 0x32, 0x04, 0x08, 0x0b, 0x10, 0x01, 0x32, 0x04, 0x08, 0x0a, 0x10, 0x01, 0x32, 0x04, 0x08, 0x04, 0x10, 0x01, 0x32, 0x04, 0x08, 0x03, 0x10, 0x01, 0x38, 0x01, 0x40, 0x04, 0x4a, 0x81, 0x09, 0x08, 0x01, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x31, 0x36, 0x10, 0x0b, 0x18, 0x04, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x31, 0x38, 0x10, 0x08, 0x18, 0x11, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x30, 0x10, 0x0d, 0x18, 0x01, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x32, 0x10, 0x0c, 0x18, 0x08, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x34, 0x10, 0x06, 0x18, 0x08, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x39, 0x10, 0x17, 0x18, 0x0b, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x36, 0x10, 0x16, 0x18, 0x08, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x31, 0x10, 0x0b, 0x18, 0x18, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x34, 0x10, 0x06, 0x18, 0x18, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x38, 0x10, 0x15, 0x18, 0x16, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x31, 0x32, 0x10, 0x04, 0x18, 0x18, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x30, 0x30, 0x32, 0x10, 0x00, 0x18, 0x00, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x39, 0x10, 0x0c, 0x18, 0x0a, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x31, 0x10, 0x01, 0x18, 0x11, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x39, 0x10, 0x17, 0x18, 0x11, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x33, 0x10, 0x04, 0x18, 0x02, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x2b, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x33, 0x36, 0x10, 0x12, 0x18, 0x0f, 0x20, 0x01, 0x2a, 0x0b, 0x0a, 0x05, 0x39, 0x38, 0x32, 0x30, 0x31, 0x10, 0x02, 0x18, 0x03, 0x2a, 0x0b, 0x0a, 0x05, 0x39, 0x38, 0x32, 0x30, 0x32, 0x10, 0x02, 0x18, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x35, 0x10, 0x0f, 0x18, 0x0d, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x13, 0x0a, 0x05, 0x39, 0x38, 0x32, 0x30, 0x32, 0x10, 0x14, 0x18, 0x11, 0x20, 0x01, 0x30, 0xd8, 0xfe, 0x05, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x37, 0x10, 0x03, 0x18, 0x0f, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x34, 0x30, 0x10, 0x14, 0x18, 0x10, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x38, 0x10, 0x12, 0x18, 0x00, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x33, 0x38, 0x10, 0x0a, 0x18, 0x02, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x31, 0x31, 0x10, 0x0f, 0x18, 0x14, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x32, 0x10, 0x00, 0x18, 0x18, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x31, 0x32, 0x10, 0x02, 0x18, 0x15, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x34, 0x10, 0x18, 0x18, 0x14, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x31, 0x35, 0x10, 0x0f, 0x18, 0x03, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x35, 0x10, 0x03, 0x18, 0x0b, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x31, 0x37, 0x10, 0x08, 0x18, 0x14, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x34, 0x10, 0x10, 0x18, 0x18, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x31, 0x39, 0x10, 0x08, 0x18, 0x0e, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x37, 0x10, 0x0b, 0x18, 0x13, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x31, 0x10, 0x0b, 0x18, 0x0a, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x34, 0x10, 0x18, 0x18, 0x10, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x33, 0x10, 0x01, 0x18, 0x08, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x33, 0x37, 0x10, 0x12, 0x18, 0x07, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x34, 0x10, 0x14, 0x18, 0x18, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x39, 0x10, 0x15, 0x18, 0x0a, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x34, 0x10, 0x18, 0x18, 0x06, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x34, 0x10, 0x18, 0x18, 0x02, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x34, 0x10, 0x02, 0x18, 0x18, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x30, 0x30, 0x31, 0x10, 0x00, 0x18, 0x00, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x31, 0x33, 0x10, 0x08, 0x18, 0x18, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x30, 0x32, 0x10, 0x18, 0x18, 0x00, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x39, 0x10, 0x17, 0x18, 0x0e, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x32, 0x10, 0x07, 0x18, 0x02, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x39, 0x10, 0x12, 0x18, 0x0a, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x34, 0x10, 0x01, 0x18, 0x02, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x13, 0x0a, 0x05, 0x39, 0x38, 0x32, 0x30, 0x31, 0x10, 0x14, 0x18, 0x12, 0x20, 0x01, 0x30, 0xd8, 0xfe, 0x05, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x35, 0x10, 0x01, 0x18, 0x0d, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x33, 0x39, 0x10, 0x10, 0x18, 0x11, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x38, 0x10, 0x00, 0x18, 0x16, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x34, 0x31, 0x10, 0x15, 0x18, 0x02, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x30, 0x38, 0x10, 0x00, 0x18, 0x00, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x33, 0x31, 0x34, 0x10, 0x18, 0x18, 0x04, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x31, 0x32, 0x10, 0x04, 0x18, 0x0d, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x32, 0x39, 0x10, 0x17, 0x18, 0x14, 0x20, 0x02, 0x30, 0x00, 0x38, 0x00, 0x12, 0x11, 0x0a, 0x05, 0x39, 0x38, 0x31, 0x31, 0x34, 0x10, 0x01, 0x18, 0x0c, 0x20, 0x01, 0x30, 0x00, 0x38, 0x00, 0x50, 0x9c, 0xef, 0xd1, 0xab, 0x06, 0x58, 0xf0, 0x9c, 0x06, 0x60, 0x00, 0x68, 0xaa, 0xa3, 0xcf, 0xa6, 0x06, 0x72, 0x07, 0x72, 0x6f, 0x6e, 0x70, 0x69, 0x63, 0x68}
	proto.Unmarshal(data, &validSC19001)
}