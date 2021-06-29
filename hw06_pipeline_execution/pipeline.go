package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, terminate In, stages ...Stage) Out {
	out := in
	for _, fun := range stages {
		newIn := make(Bi)
		func(terminate <-chan interface{}, newIn Bi, out Out) {
			go func() {
				defer close(newIn)
				for {
					select {
					case <-terminate:
						return
					case v, ok := <-out:
						if !ok {
							return
						}
						newIn <- v
					}
				}
			}()
		}(terminate, newIn, out)
		out = fun(newIn)
	}
	return out
}
