module github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/parser/tests

go 1.23.5

require (
	github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/parser v0.0.0
	github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/caching v0.0.0 // indirect
)

replace github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/parser => ../

replace github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/caching => ../../caching
