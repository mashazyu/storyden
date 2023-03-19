// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// AccountGetAvatarParams is parameters of AccountGetAvatar operation.
type AccountGetAvatarParams struct {
	// Account handle.
	AccountHandle AccountHandle
}

func unpackAccountGetAvatarParams(packed middleware.Parameters) (params AccountGetAvatarParams) {
	{
		key := middleware.ParameterKey{
			Name: "account_handle",
			In:   "path",
		}
		params.AccountHandle = packed[key].(AccountHandle)
	}
	return params
}

func decodeAccountGetAvatarParams(args [1]string, r *http.Request) (params AccountGetAvatarParams, _ error) {
	// Decode path: account_handle.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "account_handle",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotAccountHandleVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotAccountHandleVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.AccountHandle = AccountHandle(paramsDotAccountHandleVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "account_handle",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// OAuthProviderCallbackParams is parameters of OAuthProviderCallback operation.
type OAuthProviderCallbackParams struct {
	// The identifier for an OAuth2 provider such as "twitter".
	OAuthProvider string
}

func unpackOAuthProviderCallbackParams(packed middleware.Parameters) (params OAuthProviderCallbackParams) {
	{
		key := middleware.ParameterKey{
			Name: "oauth_provider",
			In:   "path",
		}
		params.OAuthProvider = packed[key].(string)
	}
	return params
}

func decodeOAuthProviderCallbackParams(args [1]string, r *http.Request) (params OAuthProviderCallbackParams, _ error) {
	// Decode path: oauth_provider.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "oauth_provider",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.OAuthProvider = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "oauth_provider",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PostsCreateParams is parameters of PostsCreate operation.
type PostsCreateParams struct {
	// Thread unique and permanent identifier.
	ThreadMark ThreadMark
}

func unpackPostsCreateParams(packed middleware.Parameters) (params PostsCreateParams) {
	{
		key := middleware.ParameterKey{
			Name: "thread_mark",
			In:   "path",
		}
		params.ThreadMark = packed[key].(ThreadMark)
	}
	return params
}

func decodePostsCreateParams(args [1]string, r *http.Request) (params PostsCreateParams, _ error) {
	// Decode path: thread_mark.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "thread_mark",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotThreadMarkVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotThreadMarkVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ThreadMark = ThreadMark(paramsDotThreadMarkVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "thread_mark",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ProfileGetParams is parameters of ProfileGet operation.
type ProfileGetParams struct {
	// Account handle.
	AccountHandle AccountHandle
}

func unpackProfileGetParams(packed middleware.Parameters) (params ProfileGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "account_handle",
			In:   "path",
		}
		params.AccountHandle = packed[key].(AccountHandle)
	}
	return params
}

func decodeProfileGetParams(args [1]string, r *http.Request) (params ProfileGetParams, _ error) {
	// Decode path: account_handle.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "account_handle",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotAccountHandleVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotAccountHandleVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.AccountHandle = AccountHandle(paramsDotAccountHandleVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "account_handle",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ThreadsGetParams is parameters of ThreadsGet operation.
type ThreadsGetParams struct {
	// Thread unique and permanent identifier.
	ThreadMark ThreadMark
}

func unpackThreadsGetParams(packed middleware.Parameters) (params ThreadsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "thread_mark",
			In:   "path",
		}
		params.ThreadMark = packed[key].(ThreadMark)
	}
	return params
}

func decodeThreadsGetParams(args [1]string, r *http.Request) (params ThreadsGetParams, _ error) {
	// Decode path: thread_mark.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "thread_mark",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotThreadMarkVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotThreadMarkVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ThreadMark = ThreadMark(paramsDotThreadMarkVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "thread_mark",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ThreadsListParams is parameters of ThreadsList operation.
type ThreadsListParams struct {
	// Show only results creeated by this user.
	Author OptAccountHandle
	// Show only results with these tags.
	Tags *TagListIDs
}

func unpackThreadsListParams(packed middleware.Parameters) (params ThreadsListParams) {
	{
		key := middleware.ParameterKey{
			Name: "author",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Author = v.(OptAccountHandle)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "tags",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Tags = v.(*TagListIDs)
		}
	}
	return params
}

func decodeThreadsListParams(args [0]string, r *http.Request) (params ThreadsListParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: author.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "author",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotAuthorVal AccountHandle
				if err := func() error {
					var paramsDotAuthorValVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotAuthorValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotAuthorVal = AccountHandle(paramsDotAuthorValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.Author.SetTo(paramsDotAuthorVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "author",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: tags.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "tags",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTagsVal TagListIDs
				if err := func() error {
					var paramsDotTagsValVal []Identifier
					if err := func() error {
						return d.DecodeArray(func(d uri.Decoder) error {
							var paramsDotTagsValValVal Identifier
							if err := func() error {
								var paramsDotTagsValValValVal string
								if err := func() error {
									val, err := d.DecodeValue()
									if err != nil {
										return err
									}

									c, err := conv.ToString(val)
									if err != nil {
										return err
									}

									paramsDotTagsValValValVal = c
									return nil
								}(); err != nil {
									return err
								}
								paramsDotTagsValValVal = Identifier(paramsDotTagsValValValVal)
								return nil
							}(); err != nil {
								return err
							}
							paramsDotTagsValVal = append(paramsDotTagsValVal, paramsDotTagsValValVal)
							return nil
						})
					}(); err != nil {
						return err
					}
					paramsDotTagsVal = TagListIDs(paramsDotTagsValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.Tags = &paramsDotTagsVal
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Tags == nil {
					return nil // optional
				}
				if err := func() error {
					if err := params.Tags.Validate(); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return errors.Wrap(err, "pointer")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "tags",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// WebAuthnGetAssertionParams is parameters of WebAuthnGetAssertion operation.
type WebAuthnGetAssertionParams struct {
	// Account handle.
	AccountHandle AccountHandle
}

func unpackWebAuthnGetAssertionParams(packed middleware.Parameters) (params WebAuthnGetAssertionParams) {
	{
		key := middleware.ParameterKey{
			Name: "account_handle",
			In:   "path",
		}
		params.AccountHandle = packed[key].(AccountHandle)
	}
	return params
}

func decodeWebAuthnGetAssertionParams(args [1]string, r *http.Request) (params WebAuthnGetAssertionParams, _ error) {
	// Decode path: account_handle.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "account_handle",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotAccountHandleVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotAccountHandleVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.AccountHandle = AccountHandle(paramsDotAccountHandleVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "account_handle",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// WebAuthnRequestCredentialParams is parameters of WebAuthnRequestCredential operation.
type WebAuthnRequestCredentialParams struct {
	// Account handle.
	AccountHandle AccountHandle
}

func unpackWebAuthnRequestCredentialParams(packed middleware.Parameters) (params WebAuthnRequestCredentialParams) {
	{
		key := middleware.ParameterKey{
			Name: "account_handle",
			In:   "path",
		}
		params.AccountHandle = packed[key].(AccountHandle)
	}
	return params
}

func decodeWebAuthnRequestCredentialParams(args [1]string, r *http.Request) (params WebAuthnRequestCredentialParams, _ error) {
	// Decode path: account_handle.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "account_handle",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotAccountHandleVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotAccountHandleVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.AccountHandle = AccountHandle(paramsDotAccountHandleVal)
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "account_handle",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}