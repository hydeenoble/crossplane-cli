package crossplane

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	fieldsPCClass = []string{"classRef"}
)

type PortableClass struct {
	u *unstructured.Unstructured
}

func NewPortableClass(u *unstructured.Unstructured) *PortableClass {
	return &PortableClass{u: u}
}
func (o *PortableClass) GetAge() string {
	return GetAge(o.u)
}

func (o *PortableClass) GetStatus() string {
	return "N/A"
}

func (o *PortableClass) GetObjectDetails() ObjectDetails {
	if o.u == nil {
		return ObjectDetails{}
	}
	return getObjectDetails(o.u)
}

func (o *PortableClass) GetRelated(filterByLabel func(metav1.GroupVersionKind, string, string) ([]unstructured.Unstructured, error)) ([]*unstructured.Unstructured, error) {
	related := make([]*unstructured.Unstructured, 0)
	obj := o.u.Object

	// Get class reference
	u, err := getObjRef(obj, fieldsPCClass)
	if err != nil {
		return related, err
	}

	related = append(related, u)
	return related, nil
}
