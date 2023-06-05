type Response struct {

	Status     string

	StatusCode int
                                    // io.ReadCloser Interface
  Body       io.ReadCloser    =>    type ReadCloser interface {
                                      // io.Reader Interface
                                      Reader    
                                      // io.Closer Interface  
                                      Closer
                                    }

}