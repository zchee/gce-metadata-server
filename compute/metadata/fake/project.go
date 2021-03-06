// Copyright 2022 The gce-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	"net/http"
)

// ProjectHandler holds project metadata handlers.
//
// Project metadata entries are stored under the following directory:
//
//	http://metadata.google.internal/computeMetadata/v1/project/
//
// See: https://cloud.google.com/compute/docs/metadata/default-metadata-values#project_metadata
type ProjectHandler struct{}

// ProjectAttributeMap map of porject attributes.
//
// The project attributes are stored under the following directory:
//
//	http://metadata.google.internal/computeMetadata/v1/project/attributes/
var ProjectAttributeMap = map[string]bool{
	// Disables legacy metadata server endpoints for all VMs in your project.
	//
	// Legacy endpoints are deprecated, always set disable-legacy-endpoints to true.
	"disable-legacy-endpoints": true,

	// Sets guest attributes for the project.
	//
	// Guest attributes are custom VM instance metadata values that you can use to publish infrequent status notifications, low volume data, or low frequency data.
	// These values are useful for indicating when startup scripts have finished or for providing other infrequent status notifications to other applications.
	//
	// Note: Any user or process on your VM instance can read and write to the namespaces and keys in guest-attributes metadata.
	//
	// For more information about guest attributes, see Setting and querying guest attributes.
	"enable-guest-attributes": true,

	// Enables or disables OS inventory for the project.
	//
	// Collects and stores OS details information. This includes information such as hostname, kernel version, architecture, and installed packages details.
	//
	// For more information about OS inventory, see Viewing operating system details.
	"enable-os-inventory": true,

	// Enables or disables SSH key management on your project.
	//
	// For more information about OS Login, see Setting up OS Login.
	"enable-oslogin": true,

	// If set, stores the default region that is used by the project.
	//
	// For more information about setting default regions, see Default region and zone.
	"google-compute-default-region": true,

	// If set, stores the default zone that is used by the project.
	//
	// For more information about setting default zones, see Default region and zone.
	"google-compute-default-zone": true,

	// If you are managing SSH keys using metadata, this attribute lets you configure public SSH keys that can connect to VMs in this project.
	// If there are multiple SSH keys, each key is separated by a newline character (\n). The value of the ssh-keys attribute is a string.
	//
	// Example:
	//  user1:ssh-rsa mypublickey user1@host.com\nuser2:ssh-rsa mypublickey user2@host.com
	//
	// SSH keys managed by OS Login aren't visible in metadata.
	"ssh-keys": true,

	// Deprecated: Use ssh-keys instead.
	"sshKeys": true,

	// Enable zonal DNS and global DNS for the VMs in your project.
	//
	// For more information about internal DNS names, see Configuring DNS names.
	"vmdnssetting": true,
}

// Attributes a directory of custom metadata values passed to the VMs in your project during startup or shutdown.
// These custom values can either be Google Cloud attributes or user-created metadata values.
//
// For a list of project-level Google Cloud attributes that you can set, see Project attributes.
//
// For more information about setting custom metadata, see Setting VM metadata.
func (h *ProjectHandler) Attributes(w http.ResponseWriter, r *http.Request) {}

// NumericProjectID is the numeric project ID (project number) of the instance, which is not the same as the project name that is visible in the Google Cloud console.
// This value is different from the project-id metadata entry value.
func (h *ProjectHandler) NumericProjectID(w http.ResponseWriter, r *http.Request) {}

// ProjectID is the project ID.
func (h *ProjectHandler) ProjectID(w http.ResponseWriter, r *http.Request) {}
