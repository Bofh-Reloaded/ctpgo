package data

type SimpleExecutionPath struct {
	Path []*PoolData
	Tokens []*TokenData
}

func NewSimpleExecutionPath(path []*PoolData, tokens []*TokenData) *SimpleExecutionPath {
	return &SimpleExecutionPath{
		Path: path,
		Tokens: tokens,
	}
}