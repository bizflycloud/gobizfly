// This file is part of gobizfly

package gobizfly

import (
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestSimpleStorageListBucket(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.SimpleStorageURL(simpleStoragePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "buckets": [
        {
            "name": "buckettest",
            "created_at": "2024-12-25T08:36:06.720000+00:00",
            "location": "hn",
            "size_kb": 0,
            "num_objects": 0,
            "default_storage_class": "COLD"
        }
    ]
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	buckets, err := client.CloudSimpleStorage.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, buckets, 1)
	bucket := buckets[0]
	assert.Equal(t, "buckettest", bucket.Name)
}

func TestSimpleStorageCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.SimpleStorageURL(simpleStoragePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)

		resp := `
{
   "bucket": 
      {
         "CreatedAt": "2024-12-26T04:00:10.290984Z",
         "DefaultStorageClass": "COLD",
         "Location": "hn",
         "Name": "testBucket",
         "NumObjects": 0,
         "SizeKb": 0
      },
   "message": "Tạo bucket thành công"
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	bucket, err := client.CloudSimpleStorage.Create(ctx, &BucketCreateRequest{
		Name:                "testBucket",
		Location:            "hn",
		Acl:                 "private",
		DefaultStorageClass: "COLD",
	})
	require.NoError(t, err)
	assert.Equal(t, "testBucket", bucket.Name)
	assert.Equal(t, "hn", bucket.Location)
}

func TestListWithBucketNameInfo(t *testing.T) {
	setup()
	defer teardown()
	// ?acl=acl&cors=cors&versioning=versioning&website_config=website_config
	urlT := testlib.SimpleStorageURL(simpleStoragePath) + "buckettest"
	mux.HandleFunc(urlT, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "bucket": {
        "name": "buckettest",
        "created_at": "2024-12-26T02:15:26.219812Z",
        "location": "hn",
        "size_kb": 0,
        "num_objects": 0,
        "default_storage_class": "COLD"
    },
    "acl": {
        "owner": {
            "id": "018d00e1f04d4a1a9dc72bd7447ee17d",
            "display_name": "cuongmv@bizflycloud.vn"
        },
        "grants": [
            {
                "permission": "READ",
                "grantee": {
                    "type": "Group",
                    "id": null,
                    "display_name": null,
                    "email": null,
                    "uri": "http://acs.amazonaws.com/groups/global/AllUsers"
                }
            },
            {
                "permission": "FULL_CONTROL",
                "grantee": {
                    "type": "CanonicalUser",
                    "id": "018d00e1f04d4a1a9dc72bd7447ee17d",
                    "display_name": "cuongmv@bizflycloud.vn",
                    "email": null,
                    "uri": null
                }
            }
        ]
    },
    "cors": {
        "rules": [
            {
                "allowed_origin": "http://another-origin.com",
                "allowed_methods": [
                    "POST"
                ],
                "allowed_headers": [
                    "Authorization"
                ],
                "exposed_headers": [],
                "max_age_seconds": 7200
            },
            {
                "allowed_origin": "http://ahoho.com",
                "allowed_methods": [
                    "PUT"
                ],
                "allowed_headers": [
                    "Content-Type"
                ],
                "exposed_headers": [],
                "max_age_seconds": 6000
            }
        ]
    },
    "versioning": {
        "status": "Enabled"
    },
    "website_config": {
        "website_url": "https://buckettest.hn-staging.ss-website.bfcplatform.vn",
        "index": "tttt.html",
        "error": "okokokoefe"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	paramListWithBucketName := ParamListWithBucketNameInfo{
		Acl:           "acl",
		Versioning:    "versioning",
		WebsiteConfig: "website_config",
		Cors:          "cors",
		BucketName:    "buckettest",
	}
	bucket, err := client.CloudSimpleStorage.ListWithBucketNameInfo(ctx, paramListWithBucketName)
	require.NoError(t, err)
	assert.Equal(t, "buckettest", bucket.Bucket.Name)
}

func TestSimpleStorageDelete(t *testing.T) {
	setup()
	defer teardown()

	var c cloudSimpleStorageService
	mux.HandleFunc(testlib.SimpleStorageURL(c.itemPath("testbucket")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})
	require.NoError(t, client.CloudSimpleStorage.Delete(ctx, "testbucket"))
}

func TestSimpleStorageUpdateAcl(t *testing.T) {
	setup()
	defer teardown()

	urlT := testlib.SimpleStorageURL(simpleStoragePath) + "buckettest"
	mux.HandleFunc(urlT, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
		resp := `
{
    "acl": {
        "message": "Bucket ACL đã được thay đổi thành: private",
        "owner": {
            "id": "018d00e1f04d4a1a9dc72bd7447ee17d",
            "display_name": "cuongmv@bizflycloud.vn"
        },
        "grants": [
            {
                "permission": "READ",
                "grantee": {
                    "type": "Group",
                    "id": null,
                    "display_name": null,
                    "email": null,
                    "uri": "http://acs.amazonaws.com/groups/global/AllUsers"
                }
            },
            {
                "permission": "FULL_CONTROL",
                "grantee": {
                    "type": "CanonicalUser",
                    "id": "018d00e1f04d4a1a9dc72bd7447ee17d",
                    "display_name": "cuongmv@bizflycloud.vn",
                    "email": null,
                    "uri": null
                }
            }
        ]
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	acl := "private"
	bucketName := "buckettest"
	bucketACL, err := client.CloudSimpleStorage.UpdateAcl(ctx, acl, bucketName)
	require.NoError(t, err)
	require.Equal(t, "Bucket ACL đã được thay đổi thành: private", bucketACL.Acl.Message)
}

func TestSimpleStorageUpdateVersioning(t *testing.T) {
	setup()
	defer teardown()

	urlT := testlib.SimpleStorageURL(simpleStoragePath) + "buckettest"
	mux.HandleFunc(urlT, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
		resp := `
{
    "versioning": {
        "message": "Bucket versioning đã được Bật",
        "status": "Enabled"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	bucketName := "buckettest"
	ResVersioning, err := client.CloudSimpleStorage.UpdateVersioning(ctx, true, bucketName)
	require.NoError(t, err)
	require.Equal(t, "Bucket versioning đã được Bật", ResVersioning.Message)
	require.Equal(t, "Enabled", ResVersioning.Status)
}

func TestSimpleStorageUpdateCors(t *testing.T) {
	setup()
	defer teardown()

	urlT := testlib.SimpleStorageURL(simpleStoragePath) + "buckettest"
	mux.HandleFunc(urlT, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
		resp := `
{
    "cors": {
        "message": "Bucket CORS \u0111\u00e3 \u0111\u01b0\u1ee3c c\u1eadp nh\u1eadt",
        "rules": [
            {
                "allowed_origin": "http://yyyyy.com",
                "allowed_methods": [
                    "GET"
                ],
                "allowed_headers": [],
                "exposed_headers": [],
                "max_age_seconds": 3600
            }
        ]
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	paramUpdateCors := ParamUpdateCors{
		Rules: []Rule{
			{
				AllowedOrigin:  "http://yyyyy.com",
				AllowedMethods: []string{},
				AllowedHeaders: []string{"PUT"},
				MaxAgeSeconds:  6400,
			},
		},
		BucketName: "buckettest",
	}
	bucketCors, err := client.CloudSimpleStorage.UpdateCors(ctx, &paramUpdateCors)
	require.NoError(t, err)
	require.Equal(t, "http://yyyyy.com", bucketCors.Rules[0].AllowedOrigin)
}

func TestSimpleStorageUpdateWebsiteConfig(t *testing.T) {
	setup()
	defer teardown()

	urlT := testlib.SimpleStorageURL(simpleStoragePath) + "buckettest"
	mux.HandleFunc(urlT, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPatch, r.Method)
		resp := `
{
    "website_config": {
        "message": "Cấu hình bucket website đã được cập nhật",
        "website_url": "https://buckettest.hn-staging.ss-website.bfcplatform.vn",
        "index": "index.html",
        "error": "gttgt"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	paramUpdateWebsiteConfig := ParamUpdateWebsiteConfig{
		Index:      "index.html",
		Error:      "okoekokeokoeko",
		BucketName: "buckettest",
	}
	bucketWebsite, err := client.CloudSimpleStorage.UpdateWebsiteConfig(ctx, &paramUpdateWebsiteConfig)
	require.NoError(t, err)
	require.Equal(t, "index.html", bucketWebsite.Index)
	require.Equal(t, "gttgt", bucketWebsite.Error)
}

func TestSimpleStorageKeyList(t *testing.T) {
	setup()
	defer teardown()

	urlT := testlib.SimpleStorageURL(simpleStorageKeyPath)
	mux.HandleFunc(urlT, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "keys": [
        {
            "user": "oked00e1f04d4a1a9dc72bd7447ee176:cuong01",
            "access_key": "okeJ8EG2V2PU9N1L3YGG"
        }
    ]
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	keyList, err := client.CloudSimpleStorage.SimpleStorageKey().ListAccessKey(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, "oked00e1f04d4a1a9dc72bd7447ee176:cuong01", keyList[0].User)
	require.Equal(t, "okeJ8EG2V2PU9N1L3YGG", keyList[0].AccessKey)
}

func TestSimpleStorageKeyGet(t *testing.T) {
	setup()
	defer teardown()

	urlT := testlib.SimpleStorageURL(simpleStorageKeyPath) + "/UITOU3XAR6TW1F2C3KYF"
	mux.HandleFunc(urlT, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "access_key": "okeOU3XAR6TW1F2C3KYF",
    "secret_key": "okeCGI1O8RprlGmysjudfzTU5ghLXML1CkBcwyo3"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	keyGet, err := client.CloudSimpleStorage.SimpleStorageKey().GetAccessKey(ctx, "UITOU3XAR6TW1F2C3KYF")
	require.NoError(t, err)
	require.Equal(t, "okeOU3XAR6TW1F2C3KYF", keyGet.AccessKey)
	require.Equal(t, "okeCGI1O8RprlGmysjudfzTU5ghLXML1CkBcwyo3", keyGet.SecretKey)
}

func TestSimpleStorageKeyCreate(t *testing.T) {
	setup()
	defer teardown()

	urlT := testlib.SimpleStorageURL(simpleStorageKeyPath)
	mux.HandleFunc(urlT, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
    "message": "Tạo Access Key Thành Công",
    "key": {
        "access_key": "okeQJF4UNILQHS8054GI",
        "secret_key": "okeMEoF6sg9ZyjTIEU8GJWPH0mRjGdseGDrokgD5"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})
	cr := KeyCreateRequest{
		SubuserId: "cuong01",
		AccessKey: "okeQJF4UNILQHS8054GI",
		SecretKey: "okeMEoF6sg9ZyjTIEU8GJWPH0mRjGdseGDrokgD5",
	}

	keyCreate, err := client.CloudSimpleStorage.SimpleStorageKey().CreateAccessKey(ctx, &cr)

	require.NoError(t, err)
	require.Equal(t, "okeQJF4UNILQHS8054GI", keyCreate.AccessKey)
	require.Equal(t, "okeMEoF6sg9ZyjTIEU8GJWPH0mRjGdseGDrokgD5", keyCreate.SecretKey)
}

func TestSimpleStorageKeyDelete(t *testing.T) {
	setup()
	defer teardown()

	var c cloudSimpleStorageKeyResource
	mux.HandleFunc(testlib.SimpleStorageURL(c.keyItemPath("gggg")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})
	require.NoError(t, client.CloudSimpleStorage.SimpleStorageKey().DeleteAccessKey(ctx, "gggg"))
}
