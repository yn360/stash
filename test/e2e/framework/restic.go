package framework

import (
	"github.com/appscode/go/crypto/rand"
	sapi "github.com/appscode/stash/api"
	"github.com/appscode/stash/client/clientset"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/client-go/pkg/api/v1"
)

func (f *Invocation) _restic() sapi.Restic {
	return sapi.Restic{
		TypeMeta: metav1.TypeMeta{
			APIVersion: sapi.SchemeGroupVersion.String(),
			Kind:       clientset.ResourceKindRestic,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      rand.WithUniqSuffix("stash"),
			Namespace: f.namespace,
		},
		Spec: sapi.ResticSpec{
			Selector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": f.app,
				},
			},
			FileGroups: []sapi.FileGroup{
				{
					Path: "/lib",
					RetentionPolicy: sapi.RetentionPolicy{
						KeepLastSnapshots: 5,
					},
				},
			},
			Schedule: "@every 1m",
		},
	}
}

func (f *Invocation) ResticForLocalBackend() sapi.Restic {
	r := f._restic()
	r.Spec.Backend = sapi.Backend{
		RepositorySecretName: "",
		Local: &sapi.LocalSpec{
			Path: "/repo",
			Volume: apiv1.Volume{
				Name: "repo",
				VolumeSource: apiv1.VolumeSource{
					EmptyDir: &apiv1.EmptyDirVolumeSource{},
				},
			},
		},
	}
	return r
}

func (f *Invocation) ResticForS3Backend() sapi.Restic {
	r := f._restic()
	r.Spec.Backend = sapi.Backend{
		RepositorySecretName: "",
		S3: &sapi.S3Spec{
			Endpoint: "s3.amazonaws.com",
			Bucket:   "stash-qa",
			Prefix:   f.app,
		},
	}
	return r
}

func (f *Invocation) ResticForGCSBackend() sapi.Restic {
	r := f._restic()
	r.Spec.Backend = sapi.Backend{
		RepositorySecretName: "",
		GCS: &sapi.GCSSpec{
			Location: "us-west1",
			Bucket:   "stash-qa",
			Prefix:   f.app,
		},
	}
	return r
}

func (f *Invocation) ResticForAzureBackend() sapi.Restic {
	r := f._restic()
	r.Spec.Backend = sapi.Backend{
		RepositorySecretName: "",
		Azure: &sapi.AzureSpec{
			Container: "stashqa",
			Prefix:    f.app,
		},
	}
	return r
}

func (f *Framework) CreateRestic(obj sapi.Restic) error {
	_, err := f.stashClient.Restics(obj.Namespace).Create(&obj)
	return err
}

func (f *Framework) DeleteRestic(meta metav1.ObjectMeta) error {
	return f.stashClient.Restics(meta.Namespace).Delete(meta.Name, deleteInForeground())
}

func (f *Framework) EventuallyRestic(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(func() *sapi.Restic {
		obj, err := f.stashClient.Restics(meta.Namespace).Get(meta.Name)
		Expect(err).NotTo(HaveOccurred())
		return obj
	})
}