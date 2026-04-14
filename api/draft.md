


goctl api go -api api-dsl/main.api -dir .





		val := r.Header.Get("User-Agent")
		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "User-Agent", val)
		newReq := r.WithContext(ctx)
		next(w, newReq)


