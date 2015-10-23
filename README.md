# GoXtremIO

## Overview
```GoXtremIO``` represents API bindings for Go that allow you to manage XtremIO storage platforms.  In the true nature of API bindings, it is intended that the functions available are basically a direct implementation of what is available through the API.

## API Compatibility
Currently only tested with XMS v3.

## Functions
Name | Description
---- | -----------
`GetEvents` | get a list or filtered list of events
`GetInitiator` | get a specific initiator by name or ID
`GetInitiators` | get a non-detailed list of initiators
`GetInitiatorGroup` | get a specific initiator group by name or ID
`GetInitiatorGroups` | get a non-detailed list of initiator groups
`GetIGFolder` | get a specific initiator group folder by name or ID
`GetIGFolders` | get a non-detailed list of initiator folders
`GetISCSIPortal` | get a specific ISCSI portal by name or ID
`GetISCSIPortals` | get a non-detailed list of ISCSI portals
`GetLunMap` | get a specific LUN map by name or ID
`GetLunMaps` | get a non-detailed list of LUN maps
`NewLunMap` | create a new LUN map
`DeleteLunMap` | delete a LUN map
`GetSnapshot` | get a specific snapshot by name or ID
`GetSnapshots` | get a non-detailed list of snapshots
`NewSnapshot` | create a new snapshot of a volume
`DeleteSnapshot` | delete a snapshot
`GetVolume` | get a specific volume by name or ID
`GetVolumes` | get a non-detailed list of volumes
`NewVolume` | create a new volume
`DeleteVolume` | delete a volume
`GetVolumeFolder` | get a specific volume folder by name or ID
`GetVolumeFolders` | get a non-detailed list of volume folders

## Examples
The package was written using test files so, these can be looked at for a more comprehensive view of how to implement the different functions.

Intialize a new client

	c, err := NewClient() // or NewClientWithArgs(endpoint, insecure, userName, password)
	if err != nil {
		panic(err)
	}

Get an Initiator Group by Name

    initiator, err := c.GetInitiatorGroup("", "VPLEX-ee20")
    if err != nil {
      panic(err)
    }

Create a Volume

    opts := &NewVolumeOptions{
      VolName: "testing1",
      VolSize: 1073741824,
    }
    
    result, err := c.NewVolume(opts)
    if err != nil {
      panic(err)
    }

Map a Volume

    opts := &NewLunMapOptions{
      VolID: 24,
      IgID:  4,
    }
    
    result, err := c.NewLunMap(opts)
    if err != nil {
      panic(err)
    }
    
For example usage you can see the [Rex-RAY](https://github.com/emccode/rexray) repo.  There, the ```goxtremio``` package is used to implement a ```Volume Manager``` across multiple storage platforms. This includes managing multipathing, mounts, and filesystems.

## Environment Variables
Name | Description
---- | -----------
`GOXTREMIO_ENDPOINT` | the API endpoint, ie. https://10.5.132.140/api/json
`GOXTREMIO_USERNAME` | the username
`GOXTREMIO_PASSWORD` | the password
`GOXTREMIO_INSECURE` | whether to skip SSL validation
`GOXTREMIO_DEBUG` | enables debug logging

## Contributions
Please contribute! The API bindings are not 100% complete based on the v3 API, so there is some work left to implement all features.  

Licensing
---------
Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

Support
-------
Please file bugs and issues on the Github issues page for this project. This is to help keep track and document everything related to this repo. For general discussions and further support you can join the [EMC {code} Community slack channel](http://community.emccode.com/). Lastly, for questions asked on [Stackoverflow.com](https://stackoverflow.com) please tag them with **EMC**. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
