// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = math.Mod
	_ = validate.Int{}
)

func decodeDescribeInstanceResponse(resp *http.Response) (res DescribeInstanceResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response InstanceInfo
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeCreateSyncActionResponse(resp *http.Response) (res CreateSyncActionResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &CreateSyncActionNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeDescribeBalloonConfigResponse(resp *http.Response) (res DescribeBalloonConfigResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Balloon
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutBalloonResponse(resp *http.Response) (res PutBalloonResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PutBalloonNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchBalloonResponse(resp *http.Response) (res PatchBalloonResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PatchBalloonNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeDescribeBalloonStatsResponse(resp *http.Response) (res DescribeBalloonStatsResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response BalloonStats
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchBalloonStatsIntervalResponse(resp *http.Response) (res PatchBalloonStatsIntervalResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PatchBalloonStatsIntervalNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutGuestBootSourceResponse(resp *http.Response) (res PutGuestBootSourceResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PutGuestBootSourceNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutGuestDriveByIDResponse(resp *http.Response) (res PutGuestDriveByIDResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PutGuestDriveByIDNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchGuestDriveByIDResponse(resp *http.Response) (res PatchGuestDriveByIDResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PatchGuestDriveByIDNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutLoggerResponse(resp *http.Response) (res PutLoggerResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PutLoggerNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetMachineConfigurationResponse(resp *http.Response) (res GetMachineConfigurationResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response MachineConfiguration
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutMachineConfigurationResponse(resp *http.Response) (res PutMachineConfigurationResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PutMachineConfigurationNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchMachineConfigurationResponse(resp *http.Response) (res PatchMachineConfigurationResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PatchMachineConfigurationNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutMetricsResponse(resp *http.Response) (res PutMetricsResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PutMetricsNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeMmdsGetResponse(resp *http.Response) (res MmdsGetResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response MmdsGetOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	case 404:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeMmdsPutResponse(resp *http.Response) (res MmdsPutResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &MmdsPutNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeMmdsPatchResponse(resp *http.Response) (res MmdsPatchResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &MmdsPatchNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeMmdsConfigPutResponse(resp *http.Response) (res MmdsConfigPutResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &MmdsConfigPutNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutGuestNetworkInterfaceByIDResponse(resp *http.Response) (res PutGuestNetworkInterfaceByIDResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PutGuestNetworkInterfaceByIDNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchGuestNetworkInterfaceByIDResponse(resp *http.Response) (res PatchGuestNetworkInterfaceByIDResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PatchGuestNetworkInterfaceByIDNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeCreateSnapshotResponse(resp *http.Response) (res CreateSnapshotResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &CreateSnapshotNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeLoadSnapshotResponse(resp *http.Response) (res LoadSnapshotResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &LoadSnapshotNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePatchVmResponse(resp *http.Response) (res PatchVmResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PatchVmNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetExportVmConfigResponse(resp *http.Response) (res GetExportVmConfigResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response FullVmConfiguration
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodePutGuestVsockResponse(resp *http.Response) (res PutGuestVsockResponse, err error) {
	switch resp.StatusCode {
	case 204:
		return &PutGuestVsockNoContent{}, nil
	case 400:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Error
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}
