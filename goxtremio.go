package goxtremio

import "errors"

const debug = false

var ErrIdOrNameNotSpecified = errors.New("Either an ID or Name must be specified")

//https://10.5.132.140/api/json/types/volumes
type getVolumesResp struct {
	Volumes []struct {
		Href string `json:"href"`
		Name string `json:"name"`
	} `json:"volumes"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) getVolumes() (resp *getVolumesResp, err error) {
	err = xms.query("/types/volumes", "", nil, &resp)
	return resp, err
}

type getVolumeResp struct {
	Content struct {
		UnalignedIoAlerts     string          `json:"unaligned-io-alerts"`
		SmallIoAlerts         string          `json:"small-io-alerts"`
		SmallIops             string          `json:"small-iops"`
		DestSnapList          []interface{}   `json:"dest-snap-list"`
		VolID                 []interface{}   `json:"vol-id"`
		ObjSeverity           string          `json:"obj-severity"`
		NumOfLunMappings      int             `json:"num-of-lun-mappings"`
		UnalignedRdBw         string          `json:"unaligned-rd-bw"`
		NumOfDestSnaps        int             `json:"num-of-dest-snaps"`
		Iops                  string          `json:"iops"`
		AccNumOfSmallWr       string          `json:"acc-num-of-small-wr"`
		AlignmentOffset       int             `json:"alignment-offset"`
		LbSize                int             `json:"lb-size"`
		LogicalSpaceInUse     string          `json:"logical-space-in-use"`
		UnalignedIoRatioLevel string          `json:"unaligned-io-ratio-level"`
		AccNumOfRd            string          `json:"acc-num-of-rd"`
		Index                 int             `json:"index"`
		SmallRdBw             string          `json:"small-rd-bw"`
		NaaName               string          `json:"naa-name"`
		AccSizeOfWr           string          `json:"acc-size-of-wr"`
		AccNumOfSmallRd       string          `json:"acc-num-of-small-rd"`
		UnalignedRdIops       string          `json:"unaligned-rd-iops"`
		WrLatency             string          `json:"wr-latency"`
		SnapgrpID             []interface{}   `json:"snapgrp-id"`
		AncestorVolID         []interface{}   `json:"ancestor-vol-id"`
		VaaiTpAlerts          string          `json:"vaai-tp-alerts"`
		CreationTime          string          `json:"creation-time"`
		RdBw                  string          `json:"rd-bw"`
		XmsID                 []interface{}   `json:"xms-id"`
		Compressible          string          `json:"compressible"`
		SmallWrBw             string          `json:"small-wr-bw"`
		AccNumOfUnalignedRd   string          `json:"acc-num-of-unaligned-rd"`
		LuName                string          `json:"lu-name"`
		UnalignedIops         string          `json:"unaligned-iops"`
		UnalignedBw           string          `json:"unaligned-bw"`
		Bw                    string          `json:"bw"`
		SmallIoRatioLevel     string          `json:"small-io-ratio-level"`
		LunMappingList        [][]interface{} `json:"lun-mapping-list"`
		VolSize               string          `json:"vol-size"`
		WrIops                string          `json:"wr-iops"`
		SysID                 []interface{}   `json:"sys-id"`
		AvgLatency            string          `json:"avg-latency"`
		RdLatency             string          `json:"rd-latency"`
		SmallWrIops           string          `json:"small-wr-iops"`
		SmallBw               string          `json:"small-bw"`
		Name                  string          `json:"name"`
		AccNumOfUnalignedWr   string          `json:"acc-num-of-unaligned-wr"`
		UnalignedWrIops       string          `json:"unaligned-wr-iops"`
		AccNumOfWr            string          `json:"acc-num-of-wr"`
		SmallIoRatio          string          `json:"small-io-ratio"`
		AccSizeOfRd           string          `json:"acc-size-of-rd"`
		UnalignedWrBw         string          `json:"unaligned-wr-bw"`
		SmallRdIops           string          `json:"small-rd-iops"`
		UnalignedIoRatio      string          `json:"unaligned-io-ratio"`
		RdIops                string          `json:"rd-iops"`
		WrBw                  string          `json:"wr-bw"`
	} `json:"content"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) getVolume(id string, name string) (resp *getVolumeResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("/types/volumes", id, map[string]string{"name": name}, &resp)
	return resp, err
}

//https://10.5.132.140/api/json/types/lun-maps
type getLunMapsResp struct {
	LunMaps []struct {
		Href string `json:"href"`
		Name string `json:"name"`
	} `json:"lun-maps"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) getLunMaps() (resp *getLunMapsResp, err error) {
	err = xms.query("/types/lun-maps", "", nil, &resp)
	return resp, err
}

type getLunMapResp struct {
	Content struct {
		TgName       string        `json:"tg-name"`
		IgIndex      int           `json:"ig-index"`
		XmsID        []interface{} `json:"xms-id"`
		MappingIndex int           `json:"mapping-index"`
		ObjSeverity  string        `json:"obj-severity"`
		TgIndex      int           `json:"tg-index"`
		Lun          int           `json:"lun"`
		IgName       string        `json:"ig-name"`
		VolIndex     int           `json:"vol-index"`
		VolName      string        `json:"vol-name"`
		MappingID    []interface{} `json:"mapping-id"`
	} `json:"content"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) getLunMap(id string, name string) (resp *getLunMapResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("/types/lun-maps", id, map[string]string{"name": name}, &resp)
	return resp, err
}

//https://10.5.132.140/api/json/types/initiators
type getInitiatorsResp struct {
	Initiators []struct {
		Href string `json:"href"`
		Name string `json:"name"`
	} `json:"initiators"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
}

func (xms *XMS) getInitiators() (resp *getInitiatorsResp, err error) {
	err = xms.query("/types/initiators", "", nil, &resp)
	return resp, err
}

type getInitiatorResp struct {
	Content struct {
		SmallIops                           string        `json:"small-iops"`
		WrLatency                           string        `json:"wr-latency"`
		ChapDiscoveryInitiatorPassword      string        `json:"chap-discovery-initiator-password"`
		ObjSeverity                         string        `json:"obj-severity"`
		RdBw                                string        `json:"rd-bw"`
		UnalignedRdBw                       string        `json:"unaligned-rd-bw"`
		ChapDiscoveryInitiatorUserName      string        `json:"chap-discovery-initiator-user-name"`
		ChapDiscoveryClusterUserName        string        `json:"chap-discovery-cluster-user-name"`
		Iops                                string        `json:"iops"`
		NumOfConnTars                       int           `json:"num-of-conn-tars"`
		AccNumOfSmallWr                     string        `json:"acc-num-of-small-wr"`
		ChapAuthenticationInitiatorPassword string        `json:"chap-authentication-initiator-password"`
		AccNumOfRd                          string        `json:"acc-num-of-rd"`
		Index                               int           `json:"index"`
		PortAddress                         string        `json:"port-address"`
		SmallRdBw                           string        `json:"small-rd-bw"`
		ChapAuthenticationInitiatorUserName string        `json:"chap-authentication-initiator-user-name"`
		IgID                                []interface{} `json:"ig-id"`
		SmallBw                             string        `json:"small-bw"`
		AccSizeOfWr                         string        `json:"acc-size-of-wr"`
		AccNumOfSmallRd                     string        `json:"acc-num-of-small-rd"`
		UnalignedRdIops                     string        `json:"unaligned-rd-iops"`
		ChapDiscoveryClusterPassword        string        `json:"chap-discovery-cluster-password"`
		ChapAuthenticationClusterPassword   string        `json:"chap-authentication-cluster-password"`
		XmsID                               []interface{} `json:"xms-id"`
		UnalignedWrIops                     string        `json:"unaligned-wr-iops"`
		AccNumOfUnalignedRd                 string        `json:"acc-num-of-unaligned-rd"`
		SmallWrBw                           string        `json:"small-wr-bw"`
		UnalignedIops                       string        `json:"unaligned-iops"`
		UnalignedBw                         string        `json:"unaligned-bw"`
		Bw                                  string        `json:"bw"`
		WrIops                              string        `json:"wr-iops"`
		AvgLatency                          string        `json:"avg-latency"`
		RdLatency                           string        `json:"rd-latency"`
		SmallWrIops                         string        `json:"small-wr-iops"`
		ChapAuthenticationClusterUserName   string        `json:"chap-authentication-cluster-user-name"`
		Name                                string        `json:"name"`
		AccNumOfUnalignedWr                 string        `json:"acc-num-of-unaligned-wr"`
		AccNumOfWr                          string        `json:"acc-num-of-wr"`
		InitiatorID                         []interface{} `json:"initiator-id"`
		AccSizeOfRd                         string        `json:"acc-size-of-rd"`
		UnalignedWrBw                       string        `json:"unaligned-wr-bw"`
		SmallRdIops                         string        `json:"small-rd-iops"`
		InitiatorConnState                  string        `json:"initiator-conn-state"`
		RdIops                              string        `json:"rd-iops"`
		WrBw                                string        `json:"wr-bw"`
		PortType                            string        `json:"port-type"`
	} `json:"content"`
	Links []Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

func (xms *XMS) getInitiator(id string, name string) (resp *getInitiatorResp, err error) {
	if id == "" && name == "" {
		return nil, ErrIdOrNameNotSpecified
	}

	err = xms.query("/types/initiators", id, map[string]string{"name": name}, &resp)
	return resp, err
}
