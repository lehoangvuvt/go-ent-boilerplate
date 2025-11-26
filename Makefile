install-tools:
	go get entgo.io/ent/cmd/ent

ent-create:
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}

ent-gen:
	go generate ./ent