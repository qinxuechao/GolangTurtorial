package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type Metric struct {
	// SV -> BE
	SvTotalCamera        int64 `json:"sv_total_camera"			gorm:"sv_total_camera"`               // SV 发送的Camara总数
	SvTotalTrack         int64 `json:"sv_total_track"			gorm:"sv_total_track"`                 // SV 发送的Track总数
	SvTotalDidTrack      int64 `json:"sv_total_did_track"			gorm:"sv_total_did_track"`             // SV 发送的Did Track总数
	SvToTalDidTrackWFace int64 `json:"sv_total_did_track_w_face"	  gorm:"sv_total_did_track_w_face"` // SV 发送的Did_with_face总数
	SvTotalFidTrack      int64 `json:"sv_total_fid_track"		gorm:"sv_total_fid_track"`             // SV 发送的Fid Track总数
	SvToTalFidTrackWFace int64 `json:"sv_total_did_track_w_face"	  gorm:"sv_total_did_track_w_face"` // SV 发送的Fid_with_face总数
	// RTP -> BE
	RtpTotalStoreIn int64 `json:"rtp_total_store_in"	 gorm:"rtp_total_store_in` // RTP 发送的store_in总数
	// MV -> BE
	SvBucketSize int64 `json:"sv_bucket_size"		gorm:"sv_bucket_size"`  // SV 未处理bucket_size
	MvBucketSize int64 `json:"mv_bucket_size" 		gorm:"mv_bucket_size"` // MV 未处理bucket_size
	// BE -> IDS
	IdsTotalCall          int64   `json:"ids_total_call" 			 gorm:"ids_total_call"`          // IDS Call 总数
	IdsTotalCallMatched   int64   `json:"ids_total_call_matched" 	 gorm:"ids_total_call_matched"`  // IDS Call Match总数
	IdsTotalCallUnmatched int64   `json:"ids_total_call_unmatched    gorm:"ids_total_call_unmatched"` // IDS Call Unmatched总数
	IdsTotalCallTimeCost  float64 `json:"ids_total_call_time_cost    gorm:"ids_total_call_time_cost"` // IDS Call 总耗时
	IdsAvgCallTimeCost    float64 `json:"ids_avg_call_time_cost"	 gorm:"ids_avg_call_time_cost"`   // IDS Call 平均耗时
	//  BE -> NPP
	NppTotalCall         int64   `json:"npp_total_call 			gorm:"npp_total_call"`        // NPP Call 总数
	NppTotalCallTimeCost float64 `json:"npp_total_call_time_cost gorm:"npp_total_call_time_cost` // NPP Call 总耗时
	NppAvgCallTimeCost   float64 `json:"npp_avg_call_time_cost"  gorm:"npp_avg_call_time_cost"`  // NPP Call 平均耗时
	//  BE -> EVENT
	EvtTotalCall         int64   `json:"evt_total_call 			gorm:"evt_total_call"`        // EVT Call 总数
	EvtTotalCallTimeCost float64 `json:"evt_total_call_time_cost gorm:"evt_total_call_time_cost` // EVT Call 总耗时
	EvtAvgCallTimeCost   float64 `json:"evt_avg_call_time_cost"  gorm:"evt_avg_call_time_cost"`  // EVT Call 平均耗时
}

// MarshalBinary -
func (e *Metric) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Metric) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}

	return nil
}

func main() {
	var metric Metric
	metric.SvBucketSize = 1000
	metric.EvtTotalCallTimeCost = 1000

	metricBinary, bErr := metric.MarshalBinary()
	if bErr != nil {
		fmt.Errorf("Unable to marshal example struct due to: %s \n", bErr)
	}
	fmt.Printf("metricBinary.MarshalBinary: %v \n", string(metricBinary))


	cacheKey := "metrics"
	t := time.Now()
	cacheField := t.Format("2006-01-02")
	fmt.Println(cacheKey)
	fmt.Println(cacheField)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.87.224:31700",
		Password: "Aibee_123@!", // no password set
		DB:       1,  // use default DB
	})

	err := rdb.HSet(ctx, cacheKey, cacheField, string(metricBinary)).Err()
	if err != nil {
		panic(err)
	}

	var newExample Metric
	cacheData, cacheErr := rdb.HGet(ctx, cacheKey, cacheField).Result()
	if cacheErr != nil {
		panic(cacheErr)
	}

	if err := newExample.UnmarshalBinary([]byte(cacheData)); err != nil {
		fmt.Printf("Unable to unmarshal data into the new example struct due to: %s \n", err)
		return
	}
	fmt.Printf("Example.UnmarshalBinary: %+v \n", newExample)

	val2, err := rdb.HGet(ctx, cacheKey, "2020-03-09").Result()
	if err == redis.Nil {
		fmt.Printf("%s does not exist", cacheKey)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}