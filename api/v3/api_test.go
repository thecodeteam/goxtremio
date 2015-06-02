package apiv3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"
)

var xms *XMS

func init() {
	endpoint := os.Getenv("GOXTREMIO_ENDPOINT")
	insecure, _ := strconv.ParseBool(os.Getenv("GOXTREMIO_INSECURE"))
	username := os.Getenv("GOXTREMIO_USERNAME")
	password := os.Getenv("GOXTREMIO_PASSWORD")

	var err error
	xms, err = New(endpoint, insecure, username, password)
	if err != nil {
		panic(err)
	}
}

func TestGetVolumes(*testing.T) {

	volumes, err := xms.GetVolumes()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volumes))

}

func TestGetVolumeById(*testing.T) {

	volume, err := xms.getVolume("25", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))

}

func TestGetVolumeByName(*testing.T) {

	volume, err := xms.getVolume("", "ubuntu-vol4")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", volume))

}

func TestGetLunMaps(*testing.T) {

	lunMaps, err := xms.getLunMaps()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", lunMaps))

}

func TestGetLunMapById(*testing.T) {

	lunMap, err := xms.getLunMap("1", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", lunMap))

}

func TestGetLunMapByName(*testing.T) {

	lunMap, err := xms.getLunMap("", "24_4_1")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", lunMap))

}

func TestGetInitiators(*testing.T) {

	initiators, err := xms.getInitiators()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiators))

}

func TestGetInitiatorById(*testing.T) {

	initiator, err := xms.getInitiator("11", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))

}

func TestGetInitiatorByName(*testing.T) {

	initiator, err := xms.getInitiator("", "ubuntu1-eth1")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiator))

}

func TestGetInitiatorGroups(*testing.T) {

	initiatorGroups, err := xms.getInitiatorGroups()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroups))

}

func TestGetInitiatorGroupByID(*testing.T) {

	initiatorGroup, err := xms.getInitiatorGroup("4", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestGetInitiatorGroupByName(*testing.T) {

	initiatorGroup, err := xms.getInitiatorGroup("", "Ubuntu")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestGetVolumeFolders(*testing.T) {

	initiatorGroups, err := xms.getVolumeFolders()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroups))

}

func TestGetVolumeFolderByID(*testing.T) {

	initiatorGroup, err := xms.getVolumeFolder("7", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestGetVolumeFolderByName(*testing.T) {

	initiatorGroup, err := xms.getVolumeFolder("", "/Ubuntu")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestGetIGFolders(*testing.T) {

	initiatorGroups, err := xms.getIGFolders()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroups))

}

func TestGetIGFolderByID(*testing.T) {

	initiatorGroup, err := xms.getIGFolder("5", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestGetIGFolderByName(*testing.T) {

	initiatorGroup, err := xms.getIGFolder("", "/ubuntu")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestGetISCSIPortals(*testing.T) {

	initiatorGroups, err := xms.getISCSIPortals()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroups))

}

func TestGetISCSIPortalByID(*testing.T) {

	initiatorGroup, err := xms.getISCSIPortal("4", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestGetISCSIPortalByName(*testing.T) {

	initiatorGroup, err := xms.getISCSIPortal("", "192.168.1.64/24")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestGetEvents(*testing.T) {

	events, err := xms.getEvents("")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", events))
}

func TestGetEventsByEventCode(*testing.T) {
	events, err := xms.getEvents("")
	if err != nil {
		panic(err)
	}

	var eventsFiltered []Event
	r, _ := regexp.Compile("^Discovered Initiators for Cluster ")

	for _, event := range events.Events {
		if event.EventCode == "5000200" && r.MatchString(event.Description) {
			eventsFiltered = append(eventsFiltered, event)
		}
	}

	fmt.Println(fmt.Sprintf("%+v", eventsFiltered))
}

func TestGetSnapshots(*testing.T) {

	initiatorGroups, err := xms.getSnapshots()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroups))

}

func TestGetSnapshotByID(*testing.T) {

	initiatorGroup, err := xms.getSnapshot("25", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestGetSnapshotByName(*testing.T) {

	initiatorGroup, err := xms.getSnapshot("", "ubuntu-vol4.snap.06022015-09:58")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", initiatorGroup))

}

func TestPostVolumes(*testing.T) {
	req := &postVolumesReq{
		//   AlignmentOffset:,
		//   LbSize:512,
		//   SysID:,
		VolName: "testing1",
		VolSize: 1073741824,
		//   ParentFolderID,
	}
	postVolumesResp, err := xms.postVolumes(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", postVolumesResp))
}

func TestDeleteVolumes(*testing.T) {
	err := xms.deleteVolumes("27", "")
	if err != nil {
		panic(err)
	}
}

func TestPostLunMaps(*testing.T) {
	req := &postLunMapsReq{
		VolID: 28,
		// VolName: "ubuntu-vol4",
		IgID: 4,
		// Lun
		// TgID
	}
	postLunMapsResp, err := xms.postLunMaps(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", postLunMapsResp))
}

func TestDeleteLunMapsByID(*testing.T) {
	err := xms.deleteLunMaps("25", "")
	if err != nil {
		panic(err)
	}

}

func TestDeleteLunMapsByName(*testing.T) {
	err := xms.deleteLunMaps("0", "testing1")
	if err != nil {
		panic(err)
	}
}

func TestPostInitiators(*testing.T) {
	req := &postInitiatorsReq{
		IgID: 4,
		// IgName
		// InitiatorName
		// PortAddress
		// InitiatorAuthenticationPassword
		// InitiatorAuthenticationUserName
		// InitiatorDiscoveryPassword
		// InitiatorDiscoveryUserName
		// InitiatorChapDiscoveryCredentials
	}

	postInitiatorsResp, err := xms.postInitiators(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", postInitiatorsResp))
}

func TestPostSnapshots(*testing.T) {
	var snapList []*SnapListItem
	snapList = append(snapList, &SnapListItem{
		AncestorVolID: "ubuntu-vol4",
		SnapVolName:   "ubuntu-vol4-snap1",
	})

	req := &postSnapshotsReq{
		SnapList: snapList,
		FolderID: "/Ubuntu",
	}

	postSnapshotsResp, err := xms.postSnapshots(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%+v", postSnapshotsResp))
}

func TestDeleteSnapshotById(*testing.T) {
	err := xms.deleteVolumes("26", "")
	if err != nil {
		panic(err)
	}
}

func TestDeleteSnapshotByName(*testing.T) {
	err := xms.deleteVolumes("", "ubuntu-vol4-snap1")
	if err != nil {
		panic(err)
	}
}
