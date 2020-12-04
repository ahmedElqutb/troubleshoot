package analyzer

func getExampleNodeContents(s string) ([]byte, error) {
	exampleNodes := `
[
  {
    "metadata": {
      "name": "biggernode-3i745",
      "selfLink": "/api/v1/nodes/biggernode-3i745",
      "uid": "6111596f-2013-4e12-a4b2-b4abfdc16e91",
      "resourceVersion": "45488694",
      "creationTimestamp": "2020-12-03T22:26:44Z",
      "labels": {
        "beta.kubernetes.io/arch": "amd64",
        "beta.kubernetes.io/instance-type": "g-8vcpu-32gb",
        "beta.kubernetes.io/os": "linux",
        "doks.digitalocean.com/node-id": "f7fadc44-4401-4d53-947c-c569adde7504",
        "doks.digitalocean.com/node-pool": "biggernode",
        "doks.digitalocean.com/node-pool-id": "af086d58-e8a1-4350-aaa0-c633d63339f2",
        "doks.digitalocean.com/version": "1.19.3-do.2",
        "failure-domain.beta.kubernetes.io/region": "sfo2",
        "kubernetes.io/arch": "amd64",
        "kubernetes.io/hostname": "biggernode-3i745",
        "kubernetes.io/os": "linux",
        "node.kubernetes.io/instance-type": "g-8vcpu-32gb",
        "region": "sfo2",
        "topology.kubernetes.io/region": "sfo2"
      },
      "annotations": {
        "alpha.kubernetes.io/provided-node-ip": "***HIDDEN***",
        "csi.volume.kubernetes.io/nodeid": "{\"dobs.csi.digitalocean.com\":\"219734582\"}",
        "io.cilium.network.ipv4-cilium-host": "***HIDDEN***",
        "io.cilium.network.ipv4-health-ip": "***HIDDEN***",
        "io.cilium.network.ipv4-pod-cidr": "***HIDDEN***/25",
        "node.alpha.kubernetes.io/ttl": "0",
        "volumes.kubernetes.io/controller-managed-attach-detach": "true"
      },
      "managedFields": [
        {
          "manager": "digitalocean-cloud-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:26:44Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:labels": {
                "f:beta.kubernetes.io/instance-type": {},
                "f:failure-domain.beta.kubernetes.io/region": {},
                "f:node.kubernetes.io/instance-type": {},
                "f:topology.kubernetes.io/region": {}
              }
            },
            "f:status": {
              "f:addresses": {
                "k:{\"type\":\"ExternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "cilium-agent",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:26:52Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:io.cilium.network.ipv4-cilium-host": {},
                "f:io.cilium.network.ipv4-health-ip": {},
                "f:io.cilium.network.ipv4-pod-cidr": {}
              }
            },
            "f:status": {
              "f:conditions": {
                "k:{\"type\":\"NetworkUnavailable\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "kube-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:33:17Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:node.alpha.kubernetes.io/ttl": {}
              },
              "f:labels": {
                "f:beta.kubernetes.io/arch": {},
                "f:beta.kubernetes.io/os": {}
              }
            },
            "f:spec": {
              "f:podCIDR": {},
              "f:podCIDRs": {
                ".": {},
                "v:\"***HIDDEN***/25\"": {}
              }
            },
            "f:status": {
              "f:volumesAttached": {}
            }
          }
        },
        {
          "manager": "kubelet",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:42:48Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                ".": {},
                "f:alpha.kubernetes.io/provided-node-ip": {},
                "f:csi.volume.kubernetes.io/nodeid": {},
                "f:volumes.kubernetes.io/controller-managed-attach-detach": {}
              },
              "f:labels": {
                ".": {},
                "f:doks.digitalocean.com/node-id": {},
                "f:doks.digitalocean.com/node-pool": {},
                "f:doks.digitalocean.com/node-pool-id": {},
                "f:doks.digitalocean.com/version": {},
                "f:kubernetes.io/arch": {},
                "f:kubernetes.io/hostname": {},
                "f:kubernetes.io/os": {},
                "f:region": {}
              }
            },
            "f:spec": {
              "f:providerID": {}
            },
            "f:status": {
              "f:addresses": {
                ".": {},
                "k:{\"type\":\"Hostname\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                },
                "k:{\"type\":\"InternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              },
              "f:allocatable": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-1Gi": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:capacity": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-1Gi": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:conditions": {
                ".": {},
                "k:{\"type\":\"DiskPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"MemoryPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"PIDPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"Ready\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              },
              "f:daemonEndpoints": {
                "f:kubeletEndpoint": {
                  "f:Port": {}
                }
              },
              "f:images": {},
              "f:nodeInfo": {
                "f:architecture": {},
                "f:bootID": {},
                "f:containerRuntimeVersion": {},
                "f:kernelVersion": {},
                "f:kubeProxyVersion": {},
                "f:kubeletVersion": {},
                "f:machineID": {},
                "f:operatingSystem": {},
                "f:osImage": {},
                "f:systemUUID": {}
              },
              "f:volumesInUse": {}
            }
          }
        }
      ]
    },
    "spec": {
      "podCIDR": "***HIDDEN***/25",
      "podCIDRs": [
        "***HIDDEN***/25"
      ],
      "providerID": "digitalocean://219734582"
    },
    "status": {
      "capacity": {
        "cpu": "8",
        "ephemeral-storage": "103176100Ki",
        "hugepages-1Gi": "0",
        "hugepages-2Mi": "0",
        "memory": "32941864Ki",
        "pods": "110"
      },
      "allocatable": {
        "cpu": "8",
        "ephemeral-storage": "95087093603",
        "hugepages-1Gi": "0",
        "hugepages-2Mi": "0",
        "memory": "29222Mi",
        "pods": "110"
      },
      "conditions": [
        {
          "type": "NetworkUnavailable",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:26:52Z",
          "lastTransitionTime": "2020-12-03T22:26:52Z",
          "reason": "CiliumIsUp",
          "message": "Cilium is running on this node"
        },
        {
          "type": "MemoryPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:48Z",
          "lastTransitionTime": "2020-12-03T22:26:43Z",
          "reason": "KubeletHasSufficientMemory",
          "message": "kubelet has sufficient memory available"
        },
        {
          "type": "DiskPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:48Z",
          "lastTransitionTime": "2020-12-03T22:26:43Z",
          "reason": "KubeletHasNoDiskPressure",
          "message": "kubelet has no disk pressure"
        },
        {
          "type": "PIDPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:48Z",
          "lastTransitionTime": "2020-12-03T22:26:43Z",
          "reason": "KubeletHasSufficientPID",
          "message": "kubelet has sufficient PID available"
        },
        {
          "type": "Ready",
          "status": "True",
          "lastHeartbeatTime": "2020-12-03T22:42:48Z",
          "lastTransitionTime": "2020-12-03T22:26:54Z",
          "reason": "KubeletReady",
          "message": "kubelet is posting ready status. AppArmor enabled"
        }
      ],
      "addresses": [
        {
          "type": "Hostname",
          "address": "biggernode-3i745"
        },
        {
          "type": "InternalIP",
          "address": "***HIDDEN***"
        },
        {
          "type": "ExternalIP",
          "address": "***HIDDEN***"
        }
      ],
      "daemonEndpoints": {
        "kubeletEndpoint": {
          "Port": 10250
        }
      },
      "nodeInfo": {
        "machineID": "f1dc61431b674c2dac8039616e53dfde",
        "systemUUID": "f1dc6143-1b67-4c2d-ac80-39616e53dfde",
        "bootID": "7d9f0bfb-2df9-4e07-a7f0-82355089195e",
        "kernelVersion": "4.19.0-11-amd64",
        "osImage": "Debian GNU/Linux 10 (buster)",
        "containerRuntimeVersion": "docker://19.3.13",
        "kubeletVersion": "v1.19.3",
        "kubeProxyVersion": "v1.19.3",
        "operatingSystem": "linux",
        "architecture": "amd64"
      },
      "images": [
        {
          "names": [
            "digitalocean/doks-debug@sha256:d1a215845d868d1d6b2a6a93cb225a892d61e131954d71b5ef45664d77d8d2c7",
            "digitalocean/doks-debug:latest"
          ],
          "sizeBytes": 752144177
        },
        {
          "names": [
            "kotsadm/kotsadm-operator@sha256:ebeb7a03864bfe794114934191411c7456f70263d77112b653a458cd904c8db4",
            "kotsadm/kotsadm-operator:v1.24.1"
          ],
          "sizeBytes": 569543593
        },
        {
          "names": [
            "cilium/cilium@sha256:3cbade4c15e7deba6b9a770475beec5e102a9f8e347c28307383a24581ebb6c4",
            "cilium/cilium:v1.8.5"
          ],
          "sizeBytes": 421109795
        },
        {
          "names": [
            "kotsadm/kotsadm@sha256:407a514c69c8cc2cd75fe6fab59f4d0b1f0fbb21d5c9ce39f7229ebb41ebd901",
            "kotsadm/kotsadm:v1.18.1"
          ],
          "sizeBytes": 371386086
        },
        {
          "names": [
            "kotsadm/kotsadm@sha256:2136d951e201ced55c8da3e11a3cb6405741245158600a95180d86cc178ca042",
            "kotsadm/kotsadm:v1.24.1"
          ],
          "sizeBytes": 331022553
        },
        {
          "names": [
            "postgres@sha256:cc8fb6b149b387fed332b5bebd144f810df544e2df514383f82f6e61698b2aea",
            "postgres:10.7"
          ],
          "sizeBytes": 229651900
        },
        {
          "names": [
            "nginx@sha256:6b1daa9462046581ac15be20277a7c75476283f969cb3a61c8725ec38d3b01c3",
            "nginx:latest"
          ],
          "sizeBytes": 132895204
        },
        {
          "names": [
            "nginx@sha256:c3a1592d2b6d275bef4087573355827b200b00ffc2d9849890a4f3aa2128c4ae",
            "nginx:1.19.4"
          ],
          "sizeBytes": 132890123
        },
        {
          "names": [
            "nginx@sha256:8aa7f6a9585d908a63e5e418dc5d14ae7467d2e36e1ab4f0d8f9d059a3d071ce",
            "nginx:1.17.7"
          ],
          "sizeBytes": 126324348
        },
        {
          "names": [
            "k8s.gcr.io/kube-apiserver@sha256:6ea8c40355df6c6c47050448e1f88cb4a5d618e9e96717818d4e11fcfe156ee0",
            "k8s.gcr.io/kube-apiserver:v1.19.3"
          ],
          "sizeBytes": 118782314
        },
        {
          "names": [
            "k8s.gcr.io/kube-proxy@sha256:1f99b26aad3a90358ad83b4065cf59002b5a913e839b70744caff4a84315a2e7",
            "k8s.gcr.io/kube-proxy:v1.19.3"
          ],
          "sizeBytes": 117686573
        },
        {
          "names": [
            "k8s.gcr.io/kube-controller-manager@sha256:1ad35b623b9123c6aab99306ba5427e2829b36b378b9b80a6e988713ac5bffd4",
            "k8s.gcr.io/kube-controller-manager:v1.19.3"
          ],
          "sizeBytes": 110811498
        },
        {
          "names": [
            "redis@sha256:82451b5a633f575c3ddd32765b228ae7d3585323dd089903bad29eefe5ac77e5",
            "redis:5"
          ],
          "sizeBytes": 98374025
        },
        {
          "names": [
            "k8s.gcr.io/autoscaling/cluster-autoscaler@sha256:d6e51b58808b5abe74b57666114efab748a6d8ba73afdb24cd2a7f91f11c4b1b",
            "k8s.gcr.io/autoscaling/cluster-autoscaler:v1.19.0"
          ],
          "sizeBytes": 90379091
        },
        {
          "names": [
            "quay.io/coreos/etcd@sha256:6dde1430abf883cde8f72ee2eea18036b508f47c2f4eda2819f795974ca34cf9",
            "quay.io/coreos/etcd:v3.4.9"
          ],
          "sizeBytes": 83761798
        },
        {
          "names": [
            "cilium/operator@sha256:487ac18d5ab4aef03f2d3436a01b1bb1858bd2635f4c7f5793eac4e7438201e3",
            "cilium/operator:v1.8.5"
          ],
          "sizeBytes": 63852651
        },
        {
          "names": [
            "digitalocean/digitalocean-cloud-controller-manager@sha256:2646651c6009e8462add1761842e43a7e5de471116aa31af5f1f8b560d13a039",
            "digitalocean/digitalocean-cloud-controller-manager:v0.1.30"
          ],
          "sizeBytes": 62940839
        },
        {
          "names": [
            "kotsadm/minio@sha256:3384c64095a62bdc41157a7527f702ce9487e2c196f74130e506ec6567f91793",
            "kotsadm/minio:v1.18.1"
          ],
          "sizeBytes": 50997780
        },
        {
          "names": [
            "quay.io/k8scsi/csi-provisioner@sha256:80755314b906a87c9d8b678de14b03605a67e66c704cfc9455bb5b38643402fc",
            "quay.io/k8scsi/csi-provisioner:v2.0.2"
          ],
          "sizeBytes": 50014788
        },
        {
          "names": [
            "fluent/fluentd@sha256:b0b9e4fa12844c68af539890e20dc83db807c6872665d5591b71fb2704a1d718",
            "fluent/fluentd:v1.7"
          ],
          "sizeBytes": 49495917
        },
        {
          "names": [
            "quay.io/k8scsi/csi-snapshotter@sha256:fe290ee144392f121a8659c6cbadd8080e4ac818459bb1b644d9b4f7fe8d8d40",
            "quay.io/k8scsi/csi-snapshotter:v3.0.0"
          ],
          "sizeBytes": 47970116
        },
        {
          "names": [
            "quay.io/k8scsi/csi-resizer@sha256:83655a67f84ac500b67fd0e7084d1eaa9828321d179269171aa0c9bc35ca8ef9",
            "quay.io/k8scsi/csi-resizer:v1.0.0"
          ],
          "sizeBytes": 47793157
        },
        {
          "names": [
            "quay.io/k8scsi/csi-attacher@sha256:f141eca4d60481b2a460cbd6720a80e40bb289b0fc7a495f5e80d88eca31d9ba",
            "quay.io/k8scsi/csi-attacher:v3.0.0"
          ],
          "sizeBytes": 47741542
        },
        {
          "names": [
            "k8s.gcr.io/kube-scheduler@sha256:54c61fbd9939006a8fe691e308d28636bffd8031af9d53a97214d6e2d27b8720",
            "k8s.gcr.io/kube-scheduler:v1.19.3"
          ],
          "sizeBytes": 45660522
        },
        {
          "names": [
            "digitalocean/dosecret-operator@sha256:52c5fc89e2ba4052f9bc0bc2676ed6b806f837ce6c50990b7829a78cff541ac0",
            "digitalocean/dosecret-operator:v1.0.2"
          ],
          "sizeBytes": 43679179
        },
        {
          "names": [
            "quay.io/k8scsi/snapshot-controller@sha256:7e4916b2e53f4bf3ae5c4c85cdaf19d7844fa651da58f9e1f07342cc5f7ec35f",
            "quay.io/k8scsi/snapshot-controller:v3.0.0"
          ],
          "sizeBytes": 43155701
        },
        {
          "names": [
            "coredns/coredns@sha256:4a6e0769130686518325b21b0c1d0688b54e7c79244d48e1b15634e98e40c6ef",
            "coredns/coredns:1.7.1"
          ],
          "sizeBytes": 42376931
        },
        {
          "names": [
            "digitalocean/kubelet-rubber-stamp@sha256:df7e2c099dfeea0702dc93dbd4bc51c999f0148d84d2b288811ee714f5921927",
            "digitalocean/kubelet-rubber-stamp:v0.1.1"
          ],
          "sizeBytes": 37250784
        },
        {
          "names": [
            "digitalocean/do-agent@sha256:ca1aa3966b919594b0dfde29bcbc5d86dbfe8f75c20337efeb45cd24baf2b084",
            "digitalocean/do-agent:3"
          ],
          "sizeBytes": 29339790
        },
        {
          "names": [
            "digitalocean/do-csi-plugin@sha256:98a6080ca1de825de135caf6a8ab8a3b8ee3423f171d46979d906d32ee97600e",
            "digitalocean/do-csi-plugin:v2.1.1"
          ],
          "sizeBytes": 26389491
        },
        {
          "names": [
            "digitalocean/doks-tokenreview@sha256:18602a5ac91d1d9067a0bf25282f55974d63ab60c4e60c5cc8c4134dde1d50fa",
            "digitalocean/doks-tokenreview:v0.0.8"
          ],
          "sizeBytes": 20843793
        },
        {
          "names": [
            "quay.io/k8scsi/csi-node-driver-registrar@sha256:a104f0f0ec5fdd007a4a85ffad95a93cfb73dd7e86296d3cc7846fde505248d3",
            "quay.io/k8scsi/csi-node-driver-registrar:v2.0.1"
          ],
          "sizeBytes": 18025755
        },
        {
          "names": [
            "curlimages/curl@sha256:cda19aa5a866b225079d5b5b3ea31895c44a5d1cf3bfefe61829af9e690e8dfe",
            "curlimages/curl:7.71.0"
          ],
          "sizeBytes": 11089089
        },
        {
          "names": [
            "busybox@sha256:4b6ad3a68d34da29bf7c8ccb5d355ba8b4babcad1f99798204e7abb43e54ee3d",
            "busybox:1.30"
          ],
          "sizeBytes": 1199418
        },
        {
          "names": [
            "cilium/cilium-init@sha256:bae879d3a3c2bb599673a01d6bfa0bbf3eeb9887083479dcfb85898479b910ad",
            "cilium/cilium-init:2019-04-05"
          ],
          "sizeBytes": 1146867
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
            "k8s.gcr.io/pause:3.1"
          ],
          "sizeBytes": 742472
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:927d98197ec1141a368550822d18fa1c60bdae27b78b0c004f705f548c07814f",
            "k8s.gcr.io/pause:3.2"
          ],
          "sizeBytes": 682696
        }
      ],
      "volumesInUse": [
        "kubernetes.io/csi/dobs.csi.digitalocean.com^f00014de-e0d8-11ea-bebe-0a58ac1420fe",
        "kubernetes.io/csi/dobs.csi.digitalocean.com^f1260e74-e0d8-11ea-b57d-0a58ac1422bd"
      ],
      "volumesAttached": [
        {
          "name": "kubernetes.io/csi/dobs.csi.digitalocean.com^f00014de-e0d8-11ea-bebe-0a58ac1420fe",
          "devicePath": ""
        },
        {
          "name": "kubernetes.io/csi/dobs.csi.digitalocean.com^f1260e74-e0d8-11ea-b57d-0a58ac1422bd",
          "devicePath": ""
        }
      ]
    }
  },
  {
    "metadata": {
      "name": "pool-yd23sqk7u-3i7i7",
      "selfLink": "/api/v1/nodes/pool-yd23sqk7u-3i7i7",
      "uid": "dd06574a-2858-4407-8fa0-2b1504d7313e",
      "resourceVersion": "45488673",
      "creationTimestamp": "2020-12-03T22:31:05Z",
      "labels": {
        "beta.kubernetes.io/arch": "amd64",
        "beta.kubernetes.io/instance-type": "s-2vcpu-4gb",
        "beta.kubernetes.io/os": "linux",
        "doks.digitalocean.com/node-id": "b8d5f546-47e1-405c-b3a6-b1981ac46d8f",
        "doks.digitalocean.com/node-pool": "pool-yd23sqk7u",
        "doks.digitalocean.com/node-pool-id": "a4b78b58-8fd0-4133-b95b-d53e3caaf28e",
        "doks.digitalocean.com/version": "1.19.3-do.2",
        "failure-domain.beta.kubernetes.io/region": "sfo2",
        "kubernetes.io/arch": "amd64",
        "kubernetes.io/hostname": "pool-yd23sqk7u-3i7i7",
        "kubernetes.io/os": "linux",
        "node.kubernetes.io/instance-type": "s-2vcpu-4gb",
        "region": "sfo2",
        "topology.kubernetes.io/region": "sfo2"
      },
      "annotations": {
        "alpha.kubernetes.io/provided-node-ip": "***HIDDEN***",
        "csi.volume.kubernetes.io/nodeid": "{\"dobs.csi.digitalocean.com\":\"219734980\"}",
        "io.cilium.network.ipv4-cilium-host": "***HIDDEN***",
        "io.cilium.network.ipv4-health-ip": "***HIDDEN***",
        "io.cilium.network.ipv4-pod-cidr": "***HIDDEN***/25",
        "node.alpha.kubernetes.io/ttl": "0",
        "volumes.kubernetes.io/controller-managed-attach-detach": "true"
      },
      "managedFields": [
        {
          "manager": "digitalocean-cloud-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:31:06Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:labels": {
                "f:beta.kubernetes.io/instance-type": {},
                "f:failure-domain.beta.kubernetes.io/region": {},
                "f:node.kubernetes.io/instance-type": {},
                "f:topology.kubernetes.io/region": {}
              }
            },
            "f:status": {
              "f:addresses": {
                "k:{\"type\":\"ExternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "cilium-agent",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:31:15Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:io.cilium.network.ipv4-cilium-host": {},
                "f:io.cilium.network.ipv4-health-ip": {},
                "f:io.cilium.network.ipv4-pod-cidr": {}
              }
            },
            "f:status": {
              "f:conditions": {
                "k:{\"type\":\"NetworkUnavailable\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "kube-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:31:15Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:node.alpha.kubernetes.io/ttl": {}
              },
              "f:labels": {
                "f:beta.kubernetes.io/arch": {},
                "f:beta.kubernetes.io/os": {}
              }
            },
            "f:spec": {
              "f:podCIDR": {},
              "f:podCIDRs": {
                ".": {},
                "v:\"***HIDDEN***/25\"": {}
              }
            }
          }
        },
        {
          "manager": "kubelet",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:42:39Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                ".": {},
                "f:alpha.kubernetes.io/provided-node-ip": {},
                "f:csi.volume.kubernetes.io/nodeid": {},
                "f:volumes.kubernetes.io/controller-managed-attach-detach": {}
              },
              "f:labels": {
                ".": {},
                "f:doks.digitalocean.com/node-id": {},
                "f:doks.digitalocean.com/node-pool": {},
                "f:doks.digitalocean.com/node-pool-id": {},
                "f:doks.digitalocean.com/version": {},
                "f:kubernetes.io/arch": {},
                "f:kubernetes.io/hostname": {},
                "f:kubernetes.io/os": {},
                "f:region": {}
              }
            },
            "f:spec": {
              "f:providerID": {}
            },
            "f:status": {
              "f:addresses": {
                ".": {},
                "k:{\"type\":\"Hostname\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                },
                "k:{\"type\":\"InternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              },
              "f:allocatable": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:capacity": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:conditions": {
                ".": {},
                "k:{\"type\":\"DiskPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"MemoryPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"PIDPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"Ready\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              },
              "f:daemonEndpoints": {
                "f:kubeletEndpoint": {
                  "f:Port": {}
                }
              },
              "f:images": {},
              "f:nodeInfo": {
                "f:architecture": {},
                "f:bootID": {},
                "f:containerRuntimeVersion": {},
                "f:kernelVersion": {},
                "f:kubeProxyVersion": {},
                "f:kubeletVersion": {},
                "f:machineID": {},
                "f:operatingSystem": {},
                "f:osImage": {},
                "f:systemUUID": {}
              }
            }
          }
        }
      ]
    },
    "spec": {
      "podCIDR": "***HIDDEN***/25",
      "podCIDRs": [
        "***HIDDEN***/25"
      ],
      "providerID": "digitalocean://219734980"
    },
    "status": {
      "capacity": {
        "cpu": "2",
        "ephemeral-storage": "82533764Ki",
        "hugepages-2Mi": "0",
        "memory": "4041568Ki",
        "pods": "110"
      },
      "allocatable": {
        "cpu": "2",
        "ephemeral-storage": "76063116777",
        "hugepages-2Mi": "0",
        "memory": "3110Mi",
        "pods": "110"
      },
      "conditions": [
        {
          "type": "NetworkUnavailable",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:31:15Z",
          "lastTransitionTime": "2020-12-03T22:31:15Z",
          "reason": "CiliumIsUp",
          "message": "Cilium is running on this node"
        },
        {
          "type": "MemoryPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:39Z",
          "lastTransitionTime": "2020-12-03T22:31:05Z",
          "reason": "KubeletHasSufficientMemory",
          "message": "kubelet has sufficient memory available"
        },
        {
          "type": "DiskPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:39Z",
          "lastTransitionTime": "2020-12-03T22:31:05Z",
          "reason": "KubeletHasNoDiskPressure",
          "message": "kubelet has no disk pressure"
        },
        {
          "type": "PIDPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:39Z",
          "lastTransitionTime": "2020-12-03T22:31:05Z",
          "reason": "KubeletHasSufficientPID",
          "message": "kubelet has sufficient PID available"
        },
        {
          "type": "Ready",
          "status": "True",
          "lastHeartbeatTime": "2020-12-03T22:42:39Z",
          "lastTransitionTime": "2020-12-03T22:31:15Z",
          "reason": "KubeletReady",
          "message": "kubelet is posting ready status. AppArmor enabled"
        }
      ],
      "addresses": [
        {
          "type": "Hostname",
          "address": "pool-yd23sqk7u-3i7i7"
        },
        {
          "type": "InternalIP",
          "address": "***HIDDEN***"
        },
        {
          "type": "ExternalIP",
          "address": "***HIDDEN***"
        }
      ],
      "daemonEndpoints": {
        "kubeletEndpoint": {
          "Port": 10250
        }
      },
      "nodeInfo": {
        "machineID": "1339576bba8440239f1ad41ac4c5c8b2",
        "systemUUID": "1339576b-ba84-4023-9f1a-d41ac4c5c8b2",
        "bootID": "f1f2cf7f-c852-410e-9574-b42a463cda3b",
        "kernelVersion": "4.19.0-11-amd64",
        "osImage": "Debian GNU/Linux 10 (buster)",
        "containerRuntimeVersion": "docker://19.3.13",
        "kubeletVersion": "v1.19.3",
        "kubeProxyVersion": "v1.19.3",
        "operatingSystem": "linux",
        "architecture": "amd64"
      },
      "images": [
        {
          "names": [
            "digitalocean/doks-debug@sha256:d1a215845d868d1d6b2a6a93cb225a892d61e131954d71b5ef45664d77d8d2c7",
            "digitalocean/doks-debug:latest"
          ],
          "sizeBytes": 752144177
        },
        {
          "names": [
            "cilium/cilium@sha256:3cbade4c15e7deba6b9a770475beec5e102a9f8e347c28307383a24581ebb6c4",
            "cilium/cilium:v1.8.5"
          ],
          "sizeBytes": 421109795
        },
        {
          "names": [
            "nginx@sha256:6b1daa9462046581ac15be20277a7c75476283f969cb3a61c8725ec38d3b01c3",
            "nginx:latest"
          ],
          "sizeBytes": 132895204
        },
        {
          "names": [
            "k8s.gcr.io/kube-apiserver@sha256:6ea8c40355df6c6c47050448e1f88cb4a5d618e9e96717818d4e11fcfe156ee0",
            "k8s.gcr.io/kube-apiserver:v1.19.3"
          ],
          "sizeBytes": 118782314
        },
        {
          "names": [
            "k8s.gcr.io/kube-proxy@sha256:1f99b26aad3a90358ad83b4065cf59002b5a913e839b70744caff4a84315a2e7",
            "k8s.gcr.io/kube-proxy:v1.19.3"
          ],
          "sizeBytes": 117686573
        },
        {
          "names": [
            "k8s.gcr.io/kube-controller-manager@sha256:1ad35b623b9123c6aab99306ba5427e2829b36b378b9b80a6e988713ac5bffd4",
            "k8s.gcr.io/kube-controller-manager:v1.19.3"
          ],
          "sizeBytes": 110811498
        },
        {
          "names": [
            "redis@sha256:82451b5a633f575c3ddd32765b228ae7d3585323dd089903bad29eefe5ac77e5",
            "redis:5"
          ],
          "sizeBytes": 98374025
        },
        {
          "names": [
            "k8s.gcr.io/autoscaling/cluster-autoscaler@sha256:d6e51b58808b5abe74b57666114efab748a6d8ba73afdb24cd2a7f91f11c4b1b",
            "k8s.gcr.io/autoscaling/cluster-autoscaler:v1.19.0"
          ],
          "sizeBytes": 90379091
        },
        {
          "names": [
            "quay.io/coreos/etcd@sha256:6dde1430abf883cde8f72ee2eea18036b508f47c2f4eda2819f795974ca34cf9",
            "quay.io/coreos/etcd:v3.4.9"
          ],
          "sizeBytes": 83761798
        },
        {
          "names": [
            "cilium/operator@sha256:487ac18d5ab4aef03f2d3436a01b1bb1858bd2635f4c7f5793eac4e7438201e3",
            "cilium/operator:v1.8.5"
          ],
          "sizeBytes": 63852651
        },
        {
          "names": [
            "digitalocean/digitalocean-cloud-controller-manager@sha256:2646651c6009e8462add1761842e43a7e5de471116aa31af5f1f8b560d13a039",
            "digitalocean/digitalocean-cloud-controller-manager:v0.1.30"
          ],
          "sizeBytes": 62940839
        },
        {
          "names": [
            "quay.io/k8scsi/csi-provisioner@sha256:80755314b906a87c9d8b678de14b03605a67e66c704cfc9455bb5b38643402fc",
            "quay.io/k8scsi/csi-provisioner:v2.0.2"
          ],
          "sizeBytes": 50014788
        },
        {
          "names": [
            "fluent/fluentd@sha256:b0b9e4fa12844c68af539890e20dc83db807c6872665d5591b71fb2704a1d718",
            "fluent/fluentd:v1.7"
          ],
          "sizeBytes": 49495917
        },
        {
          "names": [
            "quay.io/k8scsi/csi-snapshotter@sha256:fe290ee144392f121a8659c6cbadd8080e4ac818459bb1b644d9b4f7fe8d8d40",
            "quay.io/k8scsi/csi-snapshotter:v3.0.0"
          ],
          "sizeBytes": 47970116
        },
        {
          "names": [
            "quay.io/k8scsi/csi-resizer@sha256:83655a67f84ac500b67fd0e7084d1eaa9828321d179269171aa0c9bc35ca8ef9",
            "quay.io/k8scsi/csi-resizer:v1.0.0"
          ],
          "sizeBytes": 47793157
        },
        {
          "names": [
            "quay.io/k8scsi/csi-attacher@sha256:f141eca4d60481b2a460cbd6720a80e40bb289b0fc7a495f5e80d88eca31d9ba",
            "quay.io/k8scsi/csi-attacher:v3.0.0"
          ],
          "sizeBytes": 47741542
        },
        {
          "names": [
            "k8s.gcr.io/kube-scheduler@sha256:54c61fbd9939006a8fe691e308d28636bffd8031af9d53a97214d6e2d27b8720",
            "k8s.gcr.io/kube-scheduler:v1.19.3"
          ],
          "sizeBytes": 45660522
        },
        {
          "names": [
            "digitalocean/dosecret-operator@sha256:52c5fc89e2ba4052f9bc0bc2676ed6b806f837ce6c50990b7829a78cff541ac0",
            "digitalocean/dosecret-operator:v1.0.2"
          ],
          "sizeBytes": 43679179
        },
        {
          "names": [
            "quay.io/k8scsi/snapshot-controller@sha256:7e4916b2e53f4bf3ae5c4c85cdaf19d7844fa651da58f9e1f07342cc5f7ec35f",
            "quay.io/k8scsi/snapshot-controller:v3.0.0"
          ],
          "sizeBytes": 43155701
        },
        {
          "names": [
            "coredns/coredns@sha256:4a6e0769130686518325b21b0c1d0688b54e7c79244d48e1b15634e98e40c6ef",
            "coredns/coredns:1.7.1"
          ],
          "sizeBytes": 42376931
        },
        {
          "names": [
            "digitalocean/kubelet-rubber-stamp@sha256:df7e2c099dfeea0702dc93dbd4bc51c999f0148d84d2b288811ee714f5921927",
            "digitalocean/kubelet-rubber-stamp:v0.1.1"
          ],
          "sizeBytes": 37250784
        },
        {
          "names": [
            "digitalocean/do-agent@sha256:ca1aa3966b919594b0dfde29bcbc5d86dbfe8f75c20337efeb45cd24baf2b084",
            "digitalocean/do-agent:3"
          ],
          "sizeBytes": 29339790
        },
        {
          "names": [
            "digitalocean/do-csi-plugin@sha256:98a6080ca1de825de135caf6a8ab8a3b8ee3423f171d46979d906d32ee97600e",
            "digitalocean/do-csi-plugin:v2.1.1"
          ],
          "sizeBytes": 26389491
        },
        {
          "names": [
            "digitalocean/doks-tokenreview@sha256:18602a5ac91d1d9067a0bf25282f55974d63ab60c4e60c5cc8c4134dde1d50fa",
            "digitalocean/doks-tokenreview:v0.0.8"
          ],
          "sizeBytes": 20843793
        },
        {
          "names": [
            "quay.io/k8scsi/csi-node-driver-registrar@sha256:a104f0f0ec5fdd007a4a85ffad95a93cfb73dd7e86296d3cc7846fde505248d3",
            "quay.io/k8scsi/csi-node-driver-registrar:v2.0.1"
          ],
          "sizeBytes": 18025755
        },
        {
          "names": [
            "curlimages/curl@sha256:cda19aa5a866b225079d5b5b3ea31895c44a5d1cf3bfefe61829af9e690e8dfe",
            "curlimages/curl:7.71.0"
          ],
          "sizeBytes": 11089089
        },
        {
          "names": [
            "busybox@sha256:4b6ad3a68d34da29bf7c8ccb5d355ba8b4babcad1f99798204e7abb43e54ee3d",
            "busybox:1.30"
          ],
          "sizeBytes": 1199418
        },
        {
          "names": [
            "cilium/cilium-init@sha256:bae879d3a3c2bb599673a01d6bfa0bbf3eeb9887083479dcfb85898479b910ad",
            "cilium/cilium-init:2019-04-05"
          ],
          "sizeBytes": 1146867
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
            "k8s.gcr.io/pause:3.1"
          ],
          "sizeBytes": 742472
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:927d98197ec1141a368550822d18fa1c60bdae27b78b0c004f705f548c07814f",
            "k8s.gcr.io/pause:3.2"
          ],
          "sizeBytes": 682696
        }
      ]
    }
  },
  {
    "metadata": {
      "name": "pool-yd23sqk7u-3i7it",
      "selfLink": "/api/v1/nodes/pool-yd23sqk7u-3i7it",
      "uid": "bc2fdbc3-2418-4f24-a48c-0255d5a6ea27",
      "resourceVersion": "45488667",
      "creationTimestamp": "2020-12-03T22:35:33Z",
      "labels": {
        "beta.kubernetes.io/arch": "amd64",
        "beta.kubernetes.io/instance-type": "s-2vcpu-4gb",
        "beta.kubernetes.io/os": "linux",
        "doks.digitalocean.com/node-id": "f973d7f1-16ee-4ac0-87fd-b417a274822d",
        "doks.digitalocean.com/node-pool": "pool-yd23sqk7u",
        "doks.digitalocean.com/node-pool-id": "a4b78b58-8fd0-4133-b95b-d53e3caaf28e",
        "doks.digitalocean.com/version": "1.19.3-do.2",
        "failure-domain.beta.kubernetes.io/region": "sfo2",
        "kubernetes.io/arch": "amd64",
        "kubernetes.io/hostname": "pool-yd23sqk7u-3i7it",
        "kubernetes.io/os": "linux",
        "node.kubernetes.io/instance-type": "s-2vcpu-4gb",
        "region": "sfo2",
        "topology.kubernetes.io/region": "sfo2"
      },
      "annotations": {
        "alpha.kubernetes.io/provided-node-ip": "***HIDDEN***",
        "csi.volume.kubernetes.io/nodeid": "{\"dobs.csi.digitalocean.com\":\"219736094\"}",
        "io.cilium.network.ipv4-cilium-host": "***HIDDEN***",
        "io.cilium.network.ipv4-health-ip": "***HIDDEN***",
        "io.cilium.network.ipv4-pod-cidr": "***HIDDEN***/25",
        "node.alpha.kubernetes.io/ttl": "0",
        "volumes.kubernetes.io/controller-managed-attach-detach": "true"
      },
      "managedFields": [
        {
          "manager": "digitalocean-cloud-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:35:34Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:labels": {
                "f:beta.kubernetes.io/instance-type": {},
                "f:failure-domain.beta.kubernetes.io/region": {},
                "f:node.kubernetes.io/instance-type": {},
                "f:topology.kubernetes.io/region": {}
              }
            },
            "f:status": {
              "f:addresses": {
                "k:{\"type\":\"ExternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "cilium-agent",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:35:43Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:io.cilium.network.ipv4-cilium-host": {},
                "f:io.cilium.network.ipv4-health-ip": {},
                "f:io.cilium.network.ipv4-pod-cidr": {}
              }
            },
            "f:status": {
              "f:conditions": {
                "k:{\"type\":\"NetworkUnavailable\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "kube-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:35:43Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:node.alpha.kubernetes.io/ttl": {}
              },
              "f:labels": {
                "f:beta.kubernetes.io/arch": {},
                "f:beta.kubernetes.io/os": {}
              }
            },
            "f:spec": {
              "f:podCIDR": {},
              "f:podCIDRs": {
                ".": {},
                "v:\"***HIDDEN***/25\"": {}
              }
            }
          }
        },
        {
          "manager": "kubelet",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:42:37Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                ".": {},
                "f:alpha.kubernetes.io/provided-node-ip": {},
                "f:csi.volume.kubernetes.io/nodeid": {},
                "f:volumes.kubernetes.io/controller-managed-attach-detach": {}
              },
              "f:labels": {
                ".": {},
                "f:doks.digitalocean.com/node-id": {},
                "f:doks.digitalocean.com/node-pool": {},
                "f:doks.digitalocean.com/node-pool-id": {},
                "f:doks.digitalocean.com/version": {},
                "f:kubernetes.io/arch": {},
                "f:kubernetes.io/hostname": {},
                "f:kubernetes.io/os": {},
                "f:region": {}
              }
            },
            "f:spec": {
              "f:providerID": {}
            },
            "f:status": {
              "f:addresses": {
                ".": {},
                "k:{\"type\":\"Hostname\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                },
                "k:{\"type\":\"InternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              },
              "f:allocatable": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:capacity": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:conditions": {
                ".": {},
                "k:{\"type\":\"DiskPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"MemoryPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"PIDPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"Ready\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              },
              "f:daemonEndpoints": {
                "f:kubeletEndpoint": {
                  "f:Port": {}
                }
              },
              "f:images": {},
              "f:nodeInfo": {
                "f:architecture": {},
                "f:bootID": {},
                "f:containerRuntimeVersion": {},
                "f:kernelVersion": {},
                "f:kubeProxyVersion": {},
                "f:kubeletVersion": {},
                "f:machineID": {},
                "f:operatingSystem": {},
                "f:osImage": {},
                "f:systemUUID": {}
              }
            }
          }
        }
      ]
    },
    "spec": {
      "podCIDR": "***HIDDEN***/25",
      "podCIDRs": [
        "***HIDDEN***/25"
      ],
      "providerID": "digitalocean://219736094"
    },
    "status": {
      "capacity": {
        "cpu": "2",
        "ephemeral-storage": "82533764Ki",
        "hugepages-2Mi": "0",
        "memory": "4041600Ki",
        "pods": "110"
      },
      "allocatable": {
        "cpu": "2",
        "ephemeral-storage": "76063116777",
        "hugepages-2Mi": "0",
        "memory": "3110Mi",
        "pods": "110"
      },
      "conditions": [
        {
          "type": "NetworkUnavailable",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:35:43Z",
          "lastTransitionTime": "2020-12-03T22:35:43Z",
          "reason": "CiliumIsUp",
          "message": "Cilium is running on this node"
        },
        {
          "type": "MemoryPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:37Z",
          "lastTransitionTime": "2020-12-03T22:35:32Z",
          "reason": "KubeletHasSufficientMemory",
          "message": "kubelet has sufficient memory available"
        },
        {
          "type": "DiskPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:37Z",
          "lastTransitionTime": "2020-12-03T22:35:32Z",
          "reason": "KubeletHasNoDiskPressure",
          "message": "kubelet has no disk pressure"
        },
        {
          "type": "PIDPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:37Z",
          "lastTransitionTime": "2020-12-03T22:35:32Z",
          "reason": "KubeletHasSufficientPID",
          "message": "kubelet has sufficient PID available"
        },
        {
          "type": "Ready",
          "status": "True",
          "lastHeartbeatTime": "2020-12-03T22:42:37Z",
          "lastTransitionTime": "2020-12-03T22:35:43Z",
          "reason": "KubeletReady",
          "message": "kubelet is posting ready status. AppArmor enabled"
        }
      ],
      "addresses": [
        {
          "type": "Hostname",
          "address": "pool-yd23sqk7u-3i7it"
        },
        {
          "type": "InternalIP",
          "address": "***HIDDEN***"
        },
        {
          "type": "ExternalIP",
          "address": "***HIDDEN***"
        }
      ],
      "daemonEndpoints": {
        "kubeletEndpoint": {
          "Port": 10250
        }
      },
      "nodeInfo": {
        "machineID": "502cdcb6e2ec4abe81f79cae84aa8bf4",
        "systemUUID": "502cdcb6-e2ec-4abe-81f7-9cae84aa8bf4",
        "bootID": "0624450e-6ac3-453c-82a1-e3fb9805cb2b",
        "kernelVersion": "4.19.0-11-amd64",
        "osImage": "Debian GNU/Linux 10 (buster)",
        "containerRuntimeVersion": "docker://19.3.13",
        "kubeletVersion": "v1.19.3",
        "kubeProxyVersion": "v1.19.3",
        "operatingSystem": "linux",
        "architecture": "amd64"
      },
      "images": [
        {
          "names": [
            "digitalocean/doks-debug@sha256:d1a215845d868d1d6b2a6a93cb225a892d61e131954d71b5ef45664d77d8d2c7",
            "digitalocean/doks-debug:latest"
          ],
          "sizeBytes": 752144177
        },
        {
          "names": [
            "cilium/cilium@sha256:3cbade4c15e7deba6b9a770475beec5e102a9f8e347c28307383a24581ebb6c4",
            "cilium/cilium:v1.8.5"
          ],
          "sizeBytes": 421109795
        },
        {
          "names": [
            "nginx@sha256:6b1daa9462046581ac15be20277a7c75476283f969cb3a61c8725ec38d3b01c3",
            "nginx:latest"
          ],
          "sizeBytes": 132895204
        },
        {
          "names": [
            "nginx@sha256:c3a1592d2b6d275bef4087573355827b200b00ffc2d9849890a4f3aa2128c4ae",
            "nginx:1.19.4"
          ],
          "sizeBytes": 132890123
        },
        {
          "names": [
            "k8s.gcr.io/kube-apiserver@sha256:6ea8c40355df6c6c47050448e1f88cb4a5d618e9e96717818d4e11fcfe156ee0",
            "k8s.gcr.io/kube-apiserver:v1.19.3"
          ],
          "sizeBytes": 118782314
        },
        {
          "names": [
            "k8s.gcr.io/kube-proxy@sha256:1f99b26aad3a90358ad83b4065cf59002b5a913e839b70744caff4a84315a2e7",
            "k8s.gcr.io/kube-proxy:v1.19.3"
          ],
          "sizeBytes": 117686573
        },
        {
          "names": [
            "k8s.gcr.io/kube-controller-manager@sha256:1ad35b623b9123c6aab99306ba5427e2829b36b378b9b80a6e988713ac5bffd4",
            "k8s.gcr.io/kube-controller-manager:v1.19.3"
          ],
          "sizeBytes": 110811498
        },
        {
          "names": [
            "redis@sha256:82451b5a633f575c3ddd32765b228ae7d3585323dd089903bad29eefe5ac77e5",
            "redis:5"
          ],
          "sizeBytes": 98374025
        },
        {
          "names": [
            "k8s.gcr.io/autoscaling/cluster-autoscaler@sha256:d6e51b58808b5abe74b57666114efab748a6d8ba73afdb24cd2a7f91f11c4b1b",
            "k8s.gcr.io/autoscaling/cluster-autoscaler:v1.19.0"
          ],
          "sizeBytes": 90379091
        },
        {
          "names": [
            "quay.io/coreos/etcd@sha256:6dde1430abf883cde8f72ee2eea18036b508f47c2f4eda2819f795974ca34cf9",
            "quay.io/coreos/etcd:v3.4.9"
          ],
          "sizeBytes": 83761798
        },
        {
          "names": [
            "cilium/operator@sha256:487ac18d5ab4aef03f2d3436a01b1bb1858bd2635f4c7f5793eac4e7438201e3",
            "cilium/operator:v1.8.5"
          ],
          "sizeBytes": 63852651
        },
        {
          "names": [
            "digitalocean/digitalocean-cloud-controller-manager@sha256:2646651c6009e8462add1761842e43a7e5de471116aa31af5f1f8b560d13a039",
            "digitalocean/digitalocean-cloud-controller-manager:v0.1.30"
          ],
          "sizeBytes": 62940839
        },
        {
          "names": [
            "quay.io/k8scsi/csi-provisioner@sha256:80755314b906a87c9d8b678de14b03605a67e66c704cfc9455bb5b38643402fc",
            "quay.io/k8scsi/csi-provisioner:v2.0.2"
          ],
          "sizeBytes": 50014788
        },
        {
          "names": [
            "quay.io/k8scsi/csi-snapshotter@sha256:fe290ee144392f121a8659c6cbadd8080e4ac818459bb1b644d9b4f7fe8d8d40",
            "quay.io/k8scsi/csi-snapshotter:v3.0.0"
          ],
          "sizeBytes": 47970116
        },
        {
          "names": [
            "quay.io/k8scsi/csi-resizer@sha256:83655a67f84ac500b67fd0e7084d1eaa9828321d179269171aa0c9bc35ca8ef9",
            "quay.io/k8scsi/csi-resizer:v1.0.0"
          ],
          "sizeBytes": 47793157
        },
        {
          "names": [
            "quay.io/k8scsi/csi-attacher@sha256:f141eca4d60481b2a460cbd6720a80e40bb289b0fc7a495f5e80d88eca31d9ba",
            "quay.io/k8scsi/csi-attacher:v3.0.0"
          ],
          "sizeBytes": 47741542
        },
        {
          "names": [
            "k8s.gcr.io/kube-scheduler@sha256:54c61fbd9939006a8fe691e308d28636bffd8031af9d53a97214d6e2d27b8720",
            "k8s.gcr.io/kube-scheduler:v1.19.3"
          ],
          "sizeBytes": 45660522
        },
        {
          "names": [
            "digitalocean/dosecret-operator@sha256:52c5fc89e2ba4052f9bc0bc2676ed6b806f837ce6c50990b7829a78cff541ac0",
            "digitalocean/dosecret-operator:v1.0.2"
          ],
          "sizeBytes": 43679179
        },
        {
          "names": [
            "quay.io/k8scsi/snapshot-controller@sha256:7e4916b2e53f4bf3ae5c4c85cdaf19d7844fa651da58f9e1f07342cc5f7ec35f",
            "quay.io/k8scsi/snapshot-controller:v3.0.0"
          ],
          "sizeBytes": 43155701
        },
        {
          "names": [
            "coredns/coredns@sha256:4a6e0769130686518325b21b0c1d0688b54e7c79244d48e1b15634e98e40c6ef",
            "coredns/coredns:1.7.1"
          ],
          "sizeBytes": 42376931
        },
        {
          "names": [
            "digitalocean/kubelet-rubber-stamp@sha256:df7e2c099dfeea0702dc93dbd4bc51c999f0148d84d2b288811ee714f5921927",
            "digitalocean/kubelet-rubber-stamp:v0.1.1"
          ],
          "sizeBytes": 37250784
        },
        {
          "names": [
            "digitalocean/do-agent@sha256:ca1aa3966b919594b0dfde29bcbc5d86dbfe8f75c20337efeb45cd24baf2b084",
            "digitalocean/do-agent:3"
          ],
          "sizeBytes": 29339790
        },
        {
          "names": [
            "digitalocean/do-csi-plugin@sha256:98a6080ca1de825de135caf6a8ab8a3b8ee3423f171d46979d906d32ee97600e",
            "digitalocean/do-csi-plugin:v2.1.1"
          ],
          "sizeBytes": 26389491
        },
        {
          "names": [
            "digitalocean/doks-tokenreview@sha256:18602a5ac91d1d9067a0bf25282f55974d63ab60c4e60c5cc8c4134dde1d50fa",
            "digitalocean/doks-tokenreview:v0.0.8"
          ],
          "sizeBytes": 20843793
        },
        {
          "names": [
            "quay.io/k8scsi/csi-node-driver-registrar@sha256:a104f0f0ec5fdd007a4a85ffad95a93cfb73dd7e86296d3cc7846fde505248d3",
            "quay.io/k8scsi/csi-node-driver-registrar:v2.0.1"
          ],
          "sizeBytes": 18025755
        },
        {
          "names": [
            "curlimages/curl@sha256:cda19aa5a866b225079d5b5b3ea31895c44a5d1cf3bfefe61829af9e690e8dfe",
            "curlimages/curl:7.71.0"
          ],
          "sizeBytes": 11089089
        },
        {
          "names": [
            "busybox@sha256:4b6ad3a68d34da29bf7c8ccb5d355ba8b4babcad1f99798204e7abb43e54ee3d",
            "busybox:1.30"
          ],
          "sizeBytes": 1199418
        },
        {
          "names": [
            "cilium/cilium-init@sha256:bae879d3a3c2bb599673a01d6bfa0bbf3eeb9887083479dcfb85898479b910ad",
            "cilium/cilium-init:2019-04-05"
          ],
          "sizeBytes": 1146867
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
            "k8s.gcr.io/pause:3.1"
          ],
          "sizeBytes": 742472
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:927d98197ec1141a368550822d18fa1c60bdae27b78b0c004f705f548c07814f",
            "k8s.gcr.io/pause:3.2"
          ],
          "sizeBytes": 682696
        }
      ]
    }
  },
  {
    "metadata": {
      "name": "pool-yd23sqk7u-3i7v3",
      "selfLink": "/api/v1/nodes/pool-yd23sqk7u-3i7v3",
      "uid": "df8a4f35-7a8d-40bf-8c83-558c1f1801d2",
      "resourceVersion": "45489013",
      "creationTimestamp": "2020-12-03T22:40:06Z",
      "labels": {
        "beta.kubernetes.io/arch": "amd64",
        "beta.kubernetes.io/instance-type": "s-2vcpu-4gb",
        "beta.kubernetes.io/os": "linux",
        "doks.digitalocean.com/node-id": "42cb61d6-03ec-4738-9e83-5daa589d16ab",
        "doks.digitalocean.com/node-pool": "pool-yd23sqk7u",
        "doks.digitalocean.com/node-pool-id": "a4b78b58-8fd0-4133-b95b-d53e3caaf28e",
        "doks.digitalocean.com/version": "1.19.3-do.2",
        "failure-domain.beta.kubernetes.io/region": "sfo2",
        "kubernetes.io/arch": "amd64",
        "kubernetes.io/hostname": "pool-yd23sqk7u-3i7v3",
        "kubernetes.io/os": "linux",
        "node.kubernetes.io/instance-type": "s-2vcpu-4gb",
        "region": "sfo2",
        "topology.kubernetes.io/region": "sfo2"
      },
      "annotations": {
        "alpha.kubernetes.io/provided-node-ip": "***HIDDEN***",
        "csi.volume.kubernetes.io/nodeid": "{\"dobs.csi.digitalocean.com\":\"219736487\"}",
        "io.cilium.network.ipv4-cilium-host": "***HIDDEN***",
        "io.cilium.network.ipv4-health-ip": "***HIDDEN***",
        "io.cilium.network.ipv4-pod-cidr": "***HIDDEN***/25",
        "node.alpha.kubernetes.io/ttl": "0",
        "volumes.kubernetes.io/controller-managed-attach-detach": "true"
      },
      "managedFields": [
        {
          "manager": "digitalocean-cloud-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:40:07Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:labels": {
                "f:beta.kubernetes.io/instance-type": {},
                "f:failure-domain.beta.kubernetes.io/region": {},
                "f:node.kubernetes.io/instance-type": {},
                "f:topology.kubernetes.io/region": {}
              }
            },
            "f:status": {
              "f:addresses": {
                "k:{\"type\":\"ExternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "cilium-agent",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:40:15Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:io.cilium.network.ipv4-cilium-host": {},
                "f:io.cilium.network.ipv4-health-ip": {},
                "f:io.cilium.network.ipv4-pod-cidr": {}
              }
            },
            "f:status": {
              "f:conditions": {
                "k:{\"type\":\"NetworkUnavailable\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "kube-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:40:16Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:node.alpha.kubernetes.io/ttl": {}
              },
              "f:labels": {
                "f:beta.kubernetes.io/arch": {},
                "f:beta.kubernetes.io/os": {}
              }
            },
            "f:spec": {
              "f:podCIDR": {},
              "f:podCIDRs": {
                ".": {},
                "v:\"***HIDDEN***/25\"": {}
              }
            }
          }
        },
        {
          "manager": "kubelet",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:45:37Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                ".": {},
                "f:alpha.kubernetes.io/provided-node-ip": {},
                "f:csi.volume.kubernetes.io/nodeid": {},
                "f:volumes.kubernetes.io/controller-managed-attach-detach": {}
              },
              "f:labels": {
                ".": {},
                "f:doks.digitalocean.com/node-id": {},
                "f:doks.digitalocean.com/node-pool": {},
                "f:doks.digitalocean.com/node-pool-id": {},
                "f:doks.digitalocean.com/version": {},
                "f:kubernetes.io/arch": {},
                "f:kubernetes.io/hostname": {},
                "f:kubernetes.io/os": {},
                "f:region": {}
              }
            },
            "f:spec": {
              "f:providerID": {}
            },
            "f:status": {
              "f:addresses": {
                ".": {},
                "k:{\"type\":\"Hostname\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                },
                "k:{\"type\":\"InternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              },
              "f:allocatable": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:capacity": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:conditions": {
                ".": {},
                "k:{\"type\":\"DiskPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"MemoryPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"PIDPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"Ready\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              },
              "f:daemonEndpoints": {
                "f:kubeletEndpoint": {
                  "f:Port": {}
                }
              },
              "f:images": {},
              "f:nodeInfo": {
                "f:architecture": {},
                "f:bootID": {},
                "f:containerRuntimeVersion": {},
                "f:kernelVersion": {},
                "f:kubeProxyVersion": {},
                "f:kubeletVersion": {},
                "f:machineID": {},
                "f:operatingSystem": {},
                "f:osImage": {},
                "f:systemUUID": {}
              }
            }
          }
        }
      ]
    },
    "spec": {
      "podCIDR": "***HIDDEN***/25",
      "podCIDRs": [
        "***HIDDEN***/25"
      ],
      "providerID": "digitalocean://219736487"
    },
    "status": {
      "capacity": {
        "cpu": "2",
        "ephemeral-storage": "82533764Ki",
        "hugepages-2Mi": "0",
        "memory": "4041568Ki",
        "pods": "110"
      },
      "allocatable": {
        "cpu": "2",
        "ephemeral-storage": "76063116777",
        "hugepages-2Mi": "0",
        "memory": "3110Mi",
        "pods": "110"
      },
      "conditions": [
        {
          "type": "NetworkUnavailable",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:40:15Z",
          "lastTransitionTime": "2020-12-03T22:40:15Z",
          "reason": "CiliumIsUp",
          "message": "Cilium is running on this node"
        },
        {
          "type": "MemoryPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:45:37Z",
          "lastTransitionTime": "2020-12-03T22:40:05Z",
          "reason": "KubeletHasSufficientMemory",
          "message": "kubelet has sufficient memory available"
        },
        {
          "type": "DiskPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:45:37Z",
          "lastTransitionTime": "2020-12-03T22:40:05Z",
          "reason": "KubeletHasNoDiskPressure",
          "message": "kubelet has no disk pressure"
        },
        {
          "type": "PIDPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:45:37Z",
          "lastTransitionTime": "2020-12-03T22:40:05Z",
          "reason": "KubeletHasSufficientPID",
          "message": "kubelet has sufficient PID available"
        },
        {
          "type": "Ready",
          "status": "True",
          "lastHeartbeatTime": "2020-12-03T22:45:37Z",
          "lastTransitionTime": "2020-12-03T22:40:16Z",
          "reason": "KubeletReady",
          "message": "kubelet is posting ready status. AppArmor enabled"
        }
      ],
      "addresses": [
        {
          "type": "Hostname",
          "address": "pool-yd23sqk7u-3i7v3"
        },
        {
          "type": "InternalIP",
          "address": "***HIDDEN***"
        },
        {
          "type": "ExternalIP",
          "address": "***HIDDEN***"
        }
      ],
      "daemonEndpoints": {
        "kubeletEndpoint": {
          "Port": 10250
        }
      },
      "nodeInfo": {
        "machineID": "df961f4f45ea431daec4318c724a62ba",
        "systemUUID": "df961f4f-45ea-431d-aec4-318c724a62ba",
        "bootID": "e381312f-f0eb-4139-804d-9987c6bbc8dd",
        "kernelVersion": "4.19.0-11-amd64",
        "osImage": "Debian GNU/Linux 10 (buster)",
        "containerRuntimeVersion": "docker://19.3.13",
        "kubeletVersion": "v1.19.3",
        "kubeProxyVersion": "v1.19.3",
        "operatingSystem": "linux",
        "architecture": "amd64"
      },
      "images": [
        {
          "names": [
            "digitalocean/doks-debug@sha256:d1a215845d868d1d6b2a6a93cb225a892d61e131954d71b5ef45664d77d8d2c7",
            "digitalocean/doks-debug:latest"
          ],
          "sizeBytes": 752144177
        },
        {
          "names": [
            "cilium/cilium@sha256:3cbade4c15e7deba6b9a770475beec5e102a9f8e347c28307383a24581ebb6c4",
            "cilium/cilium:v1.8.5"
          ],
          "sizeBytes": 421109795
        },
        {
          "names": [
            "k8s.gcr.io/kube-apiserver@sha256:6ea8c40355df6c6c47050448e1f88cb4a5d618e9e96717818d4e11fcfe156ee0",
            "k8s.gcr.io/kube-apiserver:v1.19.3"
          ],
          "sizeBytes": 118782314
        },
        {
          "names": [
            "k8s.gcr.io/kube-proxy@sha256:1f99b26aad3a90358ad83b4065cf59002b5a913e839b70744caff4a84315a2e7",
            "k8s.gcr.io/kube-proxy:v1.19.3"
          ],
          "sizeBytes": 117686573
        },
        {
          "names": [
            "k8s.gcr.io/kube-controller-manager@sha256:1ad35b623b9123c6aab99306ba5427e2829b36b378b9b80a6e988713ac5bffd4",
            "k8s.gcr.io/kube-controller-manager:v1.19.3"
          ],
          "sizeBytes": 110811498
        },
        {
          "names": [
            "k8s.gcr.io/autoscaling/cluster-autoscaler@sha256:d6e51b58808b5abe74b57666114efab748a6d8ba73afdb24cd2a7f91f11c4b1b",
            "k8s.gcr.io/autoscaling/cluster-autoscaler:v1.19.0"
          ],
          "sizeBytes": 90379091
        },
        {
          "names": [
            "quay.io/coreos/etcd@sha256:6dde1430abf883cde8f72ee2eea18036b508f47c2f4eda2819f795974ca34cf9",
            "quay.io/coreos/etcd:v3.4.9"
          ],
          "sizeBytes": 83761798
        },
        {
          "names": [
            "cilium/operator@sha256:487ac18d5ab4aef03f2d3436a01b1bb1858bd2635f4c7f5793eac4e7438201e3",
            "cilium/operator:v1.8.5"
          ],
          "sizeBytes": 63852651
        },
        {
          "names": [
            "digitalocean/digitalocean-cloud-controller-manager@sha256:2646651c6009e8462add1761842e43a7e5de471116aa31af5f1f8b560d13a039",
            "digitalocean/digitalocean-cloud-controller-manager:v0.1.30"
          ],
          "sizeBytes": 62940839
        },
        {
          "names": [
            "quay.io/k8scsi/csi-provisioner@sha256:80755314b906a87c9d8b678de14b03605a67e66c704cfc9455bb5b38643402fc",
            "quay.io/k8scsi/csi-provisioner:v2.0.2"
          ],
          "sizeBytes": 50014788
        },
        {
          "names": [
            "quay.io/k8scsi/csi-snapshotter@sha256:fe290ee144392f121a8659c6cbadd8080e4ac818459bb1b644d9b4f7fe8d8d40",
            "quay.io/k8scsi/csi-snapshotter:v3.0.0"
          ],
          "sizeBytes": 47970116
        },
        {
          "names": [
            "quay.io/k8scsi/csi-resizer@sha256:83655a67f84ac500b67fd0e7084d1eaa9828321d179269171aa0c9bc35ca8ef9",
            "quay.io/k8scsi/csi-resizer:v1.0.0"
          ],
          "sizeBytes": 47793157
        },
        {
          "names": [
            "quay.io/k8scsi/csi-attacher@sha256:f141eca4d60481b2a460cbd6720a80e40bb289b0fc7a495f5e80d88eca31d9ba",
            "quay.io/k8scsi/csi-attacher:v3.0.0"
          ],
          "sizeBytes": 47741542
        },
        {
          "names": [
            "k8s.gcr.io/kube-scheduler@sha256:54c61fbd9939006a8fe691e308d28636bffd8031af9d53a97214d6e2d27b8720",
            "k8s.gcr.io/kube-scheduler:v1.19.3"
          ],
          "sizeBytes": 45660522
        },
        {
          "names": [
            "digitalocean/dosecret-operator@sha256:52c5fc89e2ba4052f9bc0bc2676ed6b806f837ce6c50990b7829a78cff541ac0",
            "digitalocean/dosecret-operator:v1.0.2"
          ],
          "sizeBytes": 43679179
        },
        {
          "names": [
            "quay.io/k8scsi/snapshot-controller@sha256:7e4916b2e53f4bf3ae5c4c85cdaf19d7844fa651da58f9e1f07342cc5f7ec35f",
            "quay.io/k8scsi/snapshot-controller:v3.0.0"
          ],
          "sizeBytes": 43155701
        },
        {
          "names": [
            "coredns/coredns@sha256:4a6e0769130686518325b21b0c1d0688b54e7c79244d48e1b15634e98e40c6ef",
            "coredns/coredns:1.7.1"
          ],
          "sizeBytes": 42376931
        },
        {
          "names": [
            "digitalocean/kubelet-rubber-stamp@sha256:df7e2c099dfeea0702dc93dbd4bc51c999f0148d84d2b288811ee714f5921927",
            "digitalocean/kubelet-rubber-stamp:v0.1.1"
          ],
          "sizeBytes": 37250784
        },
        {
          "names": [
            "digitalocean/do-agent@sha256:ca1aa3966b919594b0dfde29bcbc5d86dbfe8f75c20337efeb45cd24baf2b084",
            "digitalocean/do-agent:3"
          ],
          "sizeBytes": 29339790
        },
        {
          "names": [
            "digitalocean/do-csi-plugin@sha256:98a6080ca1de825de135caf6a8ab8a3b8ee3423f171d46979d906d32ee97600e",
            "digitalocean/do-csi-plugin:v2.1.1"
          ],
          "sizeBytes": 26389491
        },
        {
          "names": [
            "digitalocean/doks-tokenreview@sha256:18602a5ac91d1d9067a0bf25282f55974d63ab60c4e60c5cc8c4134dde1d50fa",
            "digitalocean/doks-tokenreview:v0.0.8"
          ],
          "sizeBytes": 20843793
        },
        {
          "names": [
            "quay.io/k8scsi/csi-node-driver-registrar@sha256:a104f0f0ec5fdd007a4a85ffad95a93cfb73dd7e86296d3cc7846fde505248d3",
            "quay.io/k8scsi/csi-node-driver-registrar:v2.0.1"
          ],
          "sizeBytes": 18025755
        },
        {
          "names": [
            "curlimages/curl@sha256:cda19aa5a866b225079d5b5b3ea31895c44a5d1cf3bfefe61829af9e690e8dfe",
            "curlimages/curl:7.71.0"
          ],
          "sizeBytes": 11089089
        },
        {
          "names": [
            "busybox@sha256:4b6ad3a68d34da29bf7c8ccb5d355ba8b4babcad1f99798204e7abb43e54ee3d",
            "busybox:1.30"
          ],
          "sizeBytes": 1199418
        },
        {
          "names": [
            "cilium/cilium-init@sha256:bae879d3a3c2bb599673a01d6bfa0bbf3eeb9887083479dcfb85898479b910ad",
            "cilium/cilium-init:2019-04-05"
          ],
          "sizeBytes": 1146867
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
            "k8s.gcr.io/pause:3.1"
          ],
          "sizeBytes": 742472
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:927d98197ec1141a368550822d18fa1c60bdae27b78b0c004f705f548c07814f",
            "k8s.gcr.io/pause:3.2"
          ],
          "sizeBytes": 682696
        }
      ]
    }
  },
  {
    "metadata": {
      "name": "smallnode-3i74t",
      "selfLink": "/api/v1/nodes/smallnode-3i74t",
      "uid": "0eb3a97a-5fae-40de-99a4-4fa9c76c979f",
      "resourceVersion": "45488658",
      "creationTimestamp": "2020-12-03T22:26:59Z",
      "labels": {
        "beta.kubernetes.io/arch": "amd64",
        "beta.kubernetes.io/instance-type": "s-1vcpu-2gb",
        "beta.kubernetes.io/os": "linux",
        "doks.digitalocean.com/node-id": "9fe332f9-1783-4200-94f2-c24424237fc0",
        "doks.digitalocean.com/node-pool": "smallnode",
        "doks.digitalocean.com/node-pool-id": "357601ea-4f8d-4ff5-aa84-ad036a9518ea",
        "doks.digitalocean.com/version": "1.19.3-do.2",
        "failure-domain.beta.kubernetes.io/region": "sfo2",
        "kubernetes.io/arch": "amd64",
        "kubernetes.io/hostname": "smallnode-3i74t",
        "kubernetes.io/os": "linux",
        "node.kubernetes.io/instance-type": "s-1vcpu-2gb",
        "region": "sfo2",
        "topology.kubernetes.io/region": "sfo2"
      },
      "annotations": {
        "alpha.kubernetes.io/provided-node-ip": "***HIDDEN***",
        "csi.volume.kubernetes.io/nodeid": "{\"dobs.csi.digitalocean.com\":\"219734583\"}",
        "io.cilium.network.ipv4-cilium-host": "***HIDDEN***",
        "io.cilium.network.ipv4-health-ip": "***HIDDEN***",
        "io.cilium.network.ipv4-pod-cidr": "***HIDDEN***/25",
        "node.alpha.kubernetes.io/ttl": "0",
        "volumes.kubernetes.io/controller-managed-attach-detach": "true"
      },
      "managedFields": [
        {
          "manager": "digitalocean-cloud-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:27:00Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:labels": {
                "f:beta.kubernetes.io/instance-type": {},
                "f:failure-domain.beta.kubernetes.io/region": {},
                "f:node.kubernetes.io/instance-type": {},
                "f:topology.kubernetes.io/region": {}
              }
            },
            "f:status": {
              "f:addresses": {
                "k:{\"type\":\"ExternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "cilium-agent",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:27:11Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:io.cilium.network.ipv4-cilium-host": {},
                "f:io.cilium.network.ipv4-health-ip": {},
                "f:io.cilium.network.ipv4-pod-cidr": {}
              }
            },
            "f:status": {
              "f:conditions": {
                "k:{\"type\":\"NetworkUnavailable\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              }
            }
          }
        },
        {
          "manager": "kube-controller-manager",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:27:19Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                "f:node.alpha.kubernetes.io/ttl": {}
              },
              "f:labels": {
                "f:beta.kubernetes.io/arch": {},
                "f:beta.kubernetes.io/os": {}
              }
            },
            "f:spec": {
              "f:podCIDR": {},
              "f:podCIDRs": {
                ".": {},
                "v:\"***HIDDEN***/25\"": {}
              }
            }
          }
        },
        {
          "manager": "kubelet",
          "operation": "Update",
          "apiVersion": "v1",
          "time": "2020-12-03T22:42:33Z",
          "fieldsType": "FieldsV1",
          "fieldsV1": {
            "f:metadata": {
              "f:annotations": {
                ".": {},
                "f:alpha.kubernetes.io/provided-node-ip": {},
                "f:csi.volume.kubernetes.io/nodeid": {},
                "f:volumes.kubernetes.io/controller-managed-attach-detach": {}
              },
              "f:labels": {
                ".": {},
                "f:doks.digitalocean.com/node-id": {},
                "f:doks.digitalocean.com/node-pool": {},
                "f:doks.digitalocean.com/node-pool-id": {},
                "f:doks.digitalocean.com/version": {},
                "f:kubernetes.io/arch": {},
                "f:kubernetes.io/hostname": {},
                "f:kubernetes.io/os": {},
                "f:region": {}
              }
            },
            "f:spec": {
              "f:providerID": {}
            },
            "f:status": {
              "f:addresses": {
                ".": {},
                "k:{\"type\":\"Hostname\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                },
                "k:{\"type\":\"InternalIP\"}": {
                  ".": {},
                  "f:address": {},
                  "f:type": {}
                }
              },
              "f:allocatable": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:capacity": {
                ".": {},
                "f:cpu": {},
                "f:ephemeral-storage": {},
                "f:hugepages-2Mi": {},
                "f:memory": {},
                "f:pods": {}
              },
              "f:conditions": {
                ".": {},
                "k:{\"type\":\"DiskPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"MemoryPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"PIDPressure\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                },
                "k:{\"type\":\"Ready\"}": {
                  ".": {},
                  "f:lastHeartbeatTime": {},
                  "f:lastTransitionTime": {},
                  "f:message": {},
                  "f:reason": {},
                  "f:status": {},
                  "f:type": {}
                }
              },
              "f:daemonEndpoints": {
                "f:kubeletEndpoint": {
                  "f:Port": {}
                }
              },
              "f:images": {},
              "f:nodeInfo": {
                "f:architecture": {},
                "f:bootID": {},
                "f:containerRuntimeVersion": {},
                "f:kernelVersion": {},
                "f:kubeProxyVersion": {},
                "f:kubeletVersion": {},
                "f:machineID": {},
                "f:operatingSystem": {},
                "f:osImage": {},
                "f:systemUUID": {}
              }
            }
          }
        }
      ]
    },
    "spec": {
      "podCIDR": "***HIDDEN***/25",
      "podCIDRs": [
        "***HIDDEN***/25"
      ],
      "providerID": "digitalocean://219734583"
    },
    "status": {
      "capacity": {
        "cpu": "1",
        "ephemeral-storage": "51570124Ki",
        "hugepages-2Mi": "0",
        "memory": "2043320Ki",
        "pods": "110"
      },
      "allocatable": {
        "cpu": "1",
        "ephemeral-storage": "47527026200",
        "hugepages-2Mi": "0",
        "memory": "1574Mi",
        "pods": "110"
      },
      "conditions": [
        {
          "type": "NetworkUnavailable",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:27:11Z",
          "lastTransitionTime": "2020-12-03T22:27:11Z",
          "reason": "CiliumIsUp",
          "message": "Cilium is running on this node"
        },
        {
          "type": "MemoryPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:33Z",
          "lastTransitionTime": "2020-12-03T22:26:59Z",
          "reason": "KubeletHasSufficientMemory",
          "message": "kubelet has sufficient memory available"
        },
        {
          "type": "DiskPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:33Z",
          "lastTransitionTime": "2020-12-03T22:26:59Z",
          "reason": "KubeletHasNoDiskPressure",
          "message": "kubelet has no disk pressure"
        },
        {
          "type": "PIDPressure",
          "status": "False",
          "lastHeartbeatTime": "2020-12-03T22:42:33Z",
          "lastTransitionTime": "2020-12-03T22:26:59Z",
          "reason": "KubeletHasSufficientPID",
          "message": "kubelet has sufficient PID available"
        },
        {
          "type": "Ready",
          "status": "True",
          "lastHeartbeatTime": "2020-12-03T22:42:33Z",
          "lastTransitionTime": "2020-12-03T22:27:19Z",
          "reason": "KubeletReady",
          "message": "kubelet is posting ready status. AppArmor enabled"
        }
      ],
      "addresses": [
        {
          "type": "Hostname",
          "address": "smallnode-3i74t"
        },
        {
          "type": "InternalIP",
          "address": "***HIDDEN***"
        },
        {
          "type": "ExternalIP",
          "address": "***HIDDEN***"
        }
      ],
      "daemonEndpoints": {
        "kubeletEndpoint": {
          "Port": 10250
        }
      },
      "nodeInfo": {
        "machineID": "4d7db425940b4e61822831735849c939",
        "systemUUID": "4d7db425-940b-4e61-8228-31735849c939",
        "bootID": "ee8362fc-6a1b-4f99-bcdb-8eea3380e16f",
        "kernelVersion": "4.19.0-11-amd64",
        "osImage": "Debian GNU/Linux 10 (buster)",
        "containerRuntimeVersion": "docker://19.3.13",
        "kubeletVersion": "v1.19.3",
        "kubeProxyVersion": "v1.19.3",
        "operatingSystem": "linux",
        "architecture": "amd64"
      },
      "images": [
        {
          "names": [
            "digitalocean/doks-debug@sha256:d1a215845d868d1d6b2a6a93cb225a892d61e131954d71b5ef45664d77d8d2c7",
            "digitalocean/doks-debug:latest"
          ],
          "sizeBytes": 752144177
        },
        {
          "names": [
            "cilium/cilium@sha256:3cbade4c15e7deba6b9a770475beec5e102a9f8e347c28307383a24581ebb6c4",
            "cilium/cilium:v1.8.5"
          ],
          "sizeBytes": 421109795
        },
        {
          "names": [
            "k8s.gcr.io/kube-apiserver@sha256:6ea8c40355df6c6c47050448e1f88cb4a5d618e9e96717818d4e11fcfe156ee0",
            "k8s.gcr.io/kube-apiserver:v1.19.3"
          ],
          "sizeBytes": 118782314
        },
        {
          "names": [
            "k8s.gcr.io/kube-proxy@sha256:1f99b26aad3a90358ad83b4065cf59002b5a913e839b70744caff4a84315a2e7",
            "k8s.gcr.io/kube-proxy:v1.19.3"
          ],
          "sizeBytes": 117686573
        },
        {
          "names": [
            "k8s.gcr.io/kube-controller-manager@sha256:1ad35b623b9123c6aab99306ba5427e2829b36b378b9b80a6e988713ac5bffd4",
            "k8s.gcr.io/kube-controller-manager:v1.19.3"
          ],
          "sizeBytes": 110811498
        },
        {
          "names": [
            "k8s.gcr.io/autoscaling/cluster-autoscaler@sha256:d6e51b58808b5abe74b57666114efab748a6d8ba73afdb24cd2a7f91f11c4b1b",
            "k8s.gcr.io/autoscaling/cluster-autoscaler:v1.19.0"
          ],
          "sizeBytes": 90379091
        },
        {
          "names": [
            "quay.io/coreos/etcd@sha256:6dde1430abf883cde8f72ee2eea18036b508f47c2f4eda2819f795974ca34cf9",
            "quay.io/coreos/etcd:v3.4.9"
          ],
          "sizeBytes": 83761798
        },
        {
          "names": [
            "cilium/operator@sha256:487ac18d5ab4aef03f2d3436a01b1bb1858bd2635f4c7f5793eac4e7438201e3",
            "cilium/operator:v1.8.5"
          ],
          "sizeBytes": 63852651
        },
        {
          "names": [
            "digitalocean/digitalocean-cloud-controller-manager@sha256:2646651c6009e8462add1761842e43a7e5de471116aa31af5f1f8b560d13a039",
            "digitalocean/digitalocean-cloud-controller-manager:v0.1.30"
          ],
          "sizeBytes": 62940839
        },
        {
          "names": [
            "quay.io/k8scsi/csi-provisioner@sha256:80755314b906a87c9d8b678de14b03605a67e66c704cfc9455bb5b38643402fc",
            "quay.io/k8scsi/csi-provisioner:v2.0.2"
          ],
          "sizeBytes": 50014788
        },
        {
          "names": [
            "quay.io/k8scsi/csi-snapshotter@sha256:fe290ee144392f121a8659c6cbadd8080e4ac818459bb1b644d9b4f7fe8d8d40",
            "quay.io/k8scsi/csi-snapshotter:v3.0.0"
          ],
          "sizeBytes": 47970116
        },
        {
          "names": [
            "quay.io/k8scsi/csi-resizer@sha256:83655a67f84ac500b67fd0e7084d1eaa9828321d179269171aa0c9bc35ca8ef9",
            "quay.io/k8scsi/csi-resizer:v1.0.0"
          ],
          "sizeBytes": 47793157
        },
        {
          "names": [
            "quay.io/k8scsi/csi-attacher@sha256:f141eca4d60481b2a460cbd6720a80e40bb289b0fc7a495f5e80d88eca31d9ba",
            "quay.io/k8scsi/csi-attacher:v3.0.0"
          ],
          "sizeBytes": 47741542
        },
        {
          "names": [
            "k8s.gcr.io/kube-scheduler@sha256:54c61fbd9939006a8fe691e308d28636bffd8031af9d53a97214d6e2d27b8720",
            "k8s.gcr.io/kube-scheduler:v1.19.3"
          ],
          "sizeBytes": 45660522
        },
        {
          "names": [
            "digitalocean/dosecret-operator@sha256:52c5fc89e2ba4052f9bc0bc2676ed6b806f837ce6c50990b7829a78cff541ac0",
            "digitalocean/dosecret-operator:v1.0.2"
          ],
          "sizeBytes": 43679179
        },
        {
          "names": [
            "quay.io/k8scsi/snapshot-controller@sha256:7e4916b2e53f4bf3ae5c4c85cdaf19d7844fa651da58f9e1f07342cc5f7ec35f",
            "quay.io/k8scsi/snapshot-controller:v3.0.0"
          ],
          "sizeBytes": 43155701
        },
        {
          "names": [
            "coredns/coredns@sha256:4a6e0769130686518325b21b0c1d0688b54e7c79244d48e1b15634e98e40c6ef",
            "coredns/coredns:1.7.1"
          ],
          "sizeBytes": 42376931
        },
        {
          "names": [
            "digitalocean/kubelet-rubber-stamp@sha256:df7e2c099dfeea0702dc93dbd4bc51c999f0148d84d2b288811ee714f5921927",
            "digitalocean/kubelet-rubber-stamp:v0.1.1"
          ],
          "sizeBytes": 37250784
        },
        {
          "names": [
            "digitalocean/do-agent@sha256:ca1aa3966b919594b0dfde29bcbc5d86dbfe8f75c20337efeb45cd24baf2b084",
            "digitalocean/do-agent:3"
          ],
          "sizeBytes": 29339790
        },
        {
          "names": [
            "digitalocean/do-csi-plugin@sha256:98a6080ca1de825de135caf6a8ab8a3b8ee3423f171d46979d906d32ee97600e",
            "digitalocean/do-csi-plugin:v2.1.1"
          ],
          "sizeBytes": 26389491
        },
        {
          "names": [
            "digitalocean/doks-tokenreview@sha256:18602a5ac91d1d9067a0bf25282f55974d63ab60c4e60c5cc8c4134dde1d50fa",
            "digitalocean/doks-tokenreview:v0.0.8"
          ],
          "sizeBytes": 20843793
        },
        {
          "names": [
            "quay.io/k8scsi/csi-node-driver-registrar@sha256:a104f0f0ec5fdd007a4a85ffad95a93cfb73dd7e86296d3cc7846fde505248d3",
            "quay.io/k8scsi/csi-node-driver-registrar:v2.0.1"
          ],
          "sizeBytes": 18025755
        },
        {
          "names": [
            "curlimages/curl@sha256:cda19aa5a866b225079d5b5b3ea31895c44a5d1cf3bfefe61829af9e690e8dfe",
            "curlimages/curl:7.71.0"
          ],
          "sizeBytes": 11089089
        },
        {
          "names": [
            "busybox@sha256:4b6ad3a68d34da29bf7c8ccb5d355ba8b4babcad1f99798204e7abb43e54ee3d",
            "busybox:1.30"
          ],
          "sizeBytes": 1199418
        },
        {
          "names": [
            "cilium/cilium-init@sha256:bae879d3a3c2bb599673a01d6bfa0bbf3eeb9887083479dcfb85898479b910ad",
            "cilium/cilium-init:2019-04-05"
          ],
          "sizeBytes": 1146867
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:f78411e19d84a252e53bff71a4407a5686c46983a2c2eeed83929b888179acea",
            "k8s.gcr.io/pause:3.1"
          ],
          "sizeBytes": 742472
        },
        {
          "names": [
            "k8s.gcr.io/pause@sha256:927d98197ec1141a368550822d18fa1c60bdae27b78b0c004f705f548c07814f",
            "k8s.gcr.io/pause:3.2"
          ],
          "sizeBytes": 682696
        }
      ]
    }
  }
]
`
	return []byte(exampleNodes), nil
}
