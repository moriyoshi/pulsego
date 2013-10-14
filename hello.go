package main

import (
    "github.com/moriyoshi/pulsego"
    "fmt"
)

func send(ch chan int, val int) {
    ch <- val
}

func myMain(sync_ch chan int, pa *pulsego.PulseMainLoop) {
    defer send(sync_ch, 0)
    ctx := pa.NewContext("default", 0)
    if ctx == nil {
        fmt.Println("Failed to create a new context")
        return
    }
    defer ctx.Dispose();
    st := ctx.NewStream("default", &pulsego.PulseSampleSpec {
        Format:pulsego.SAMPLE_FLOAT32LE, Rate:22500, Channels: 1 })
    if st == nil {
        fmt.Println("Failed to create a new stream")
        return
    }
    defer st.Dispose()
    st.ConnectToSink()
    var samples []float32 = make([]float32, 65536)
    period := 40
    count := 0
    v := -1
    for {
        for i := range samples {
            if count < period / 2 {
                samples[i] = -0.8
            } else {
                samples[i] = 0.8
            }
            count += 1;
            if (count >= period) {
                count = 0
            }
            if i % 40 == 0 {
                if period < 10 || period > 100 { v = -v }
                period += v
            }
        }
        st.Write(samples, pulsego.SEEK_RELATIVE);
    }
}

func main() {
    pa := pulsego.NewPulseMainLoop();
    defer pa.Dispose()
    pa.Start()

    sync_ch := make(chan int)
    go myMain(sync_ch, pa)
    <-sync_ch

}
