package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	next := in

	for _, stage := range stages {
		next = func(ch In) Out {
			stageCh := make(Bi)

			go func() {
				defer close(stageCh)

				for {
					select {
					case <-done:
						return
					case v, ok := <-ch:
						if !ok {
							return
						}

						stageCh <- v
					}
				}
			}()

			return stage(stageCh)
		}(next)
	}

	return next
}
