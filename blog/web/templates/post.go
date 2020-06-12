package template

var PostBody = `
	<h1>
		{{ .Post.Title }}
	</h1>
	<div id="content">
		{{ .Post.Content }}
	</div>
`
