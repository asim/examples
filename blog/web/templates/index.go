package template

var IndexBody = `
	{{ range $index, $post := .Posts }}
	<div class="entry" style="text-align: left">
		<h1>
			<a href="post/{{ $post.Slug }}">{{ $post.Title }}</a>
		</h1>
		<div>
			{{ range $index, $tagName := $post.TagNames }}
			<span class="tag">{{ $tagName }}</span>
			{{ end }}
		</div>
		<br />
		<div id="content">
			{{ $post.Content }}
		</div>
	</div>
	<br />
	<br />
	{{ end }}
	{{ if not .Posts }}

	<h1>No posts here yet.</h1>

	{{ end }}
`
