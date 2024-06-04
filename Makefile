gen:
	go run github.com/Khan/genqlient genqlient.yaml
	
update-schema:
	get-graphql-schema \
		-h "Authorization=Bearer ${CF_API_TOKEN}" \
		https://api.cloudflare.com/client/v4/graphql \
		> schema.graphql
