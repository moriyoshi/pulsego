package main

import "pulse"

func myMain(sync_ch chan int, pa *pulse.PulseMainLoop) {
    ctx := pa.NewContext();
    defer ctx.Dispose();
    st := ctx.NewStream(&pulse.PulseSampleSpec {
        Format:pulse.SAMPLE_FLOAT32LE, Rate:22500, Channels: 1 });
    defer st.Dispose();
    st.ConnectToSink();
    var samples [4096]float32;
    const period = 40;
    count := 0;
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
        }
        st.Write(samples, pulse.SEEK_RELATIVE);
    }
    sync_ch <- 1
}

func main() {
    pa := pulse.NewPulseMainLoop();
    defer pa.Dispose();
    pa.Start();

    sync_ch := make(chan int);
    go myMain(sync_ch, pa);
    <-sync_ch;

}
