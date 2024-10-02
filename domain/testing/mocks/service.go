// Package mocks the service domain. Service for testing purpose
package mocks

type Service struct {
	*ServiceBase // DO NOT REMOVE this line. ServiceBase implements IServiceBase interface

	/* Add your mocks use cases here. To do it follow the next code pattern:

	MethodName_ func(param1 paramType, ..., paramN paramType ) (returnType1, ..., returnTypeN)

	*/

}

/*
	Method Implementation pattern:

func(s *Service) MethodName(param1 paramType, ..., paramN paramType ) (returnType1, ..., returnTypeN) {
	return MethodName_(param1, ..., paramN)
}
*/
