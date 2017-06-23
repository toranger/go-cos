package cos

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestServiceService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `<ListAllMyBucketsResult>
	<Owner>
		<ID>xbaccxx</ID>
		<DisplayName>100000760461</DisplayName>
	</Owner>
	<Buckets>
		<Bucket>
			<Name>huadong-1253846586</Name>
			<Location>cn-east</Location>
			<CreateDate>2017-06-16T13:08:28Z</CreateDate>
		</Bucket>
		<Bucket>
			<Name>huanan-1253846586</Name>
			<Location>cn-south</Location>
			<CreateDate>2017-06-10T09:00:07Z</CreateDate>
		</Bucket>
	</Buckets>
</ListAllMyBucketsResult>`)
	})

	ref, _, err := client.Service.Get(context.Background())
	if err != nil {
		t.Fatalf("Service.Get returned error: %v", err)
	}

	want := &ServiceGetResult{
		XMLName: xml.Name{Local: "ListAllMyBucketsResult"},
		Owner: &struct {
			ID          string
			DisplayName string
		}{
			"xbaccxx",
			"100000760461",
		},
		Buckets: []struct {
			Name       string
			Location   string
			CreateDate string
		}{
			{"huadong-1253846586", "cn-east", "2017-06-16T13:08:28Z"},
			{"huanan-1253846586", "cn-south", "2017-06-10T09:00:07Z"},
		},
	}

	if !reflect.DeepEqual(ref, want) {
		t.Errorf("Service.Get returned %+v, want %+v", ref, want)
	}
}
