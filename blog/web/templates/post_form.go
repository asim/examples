package template

var NewPostBody = `
	<h1>
		New Post
	</h1>
	<div id="content" style="text-align: left;">
	<form action="api/post" method="post">
		<div class="form-group">
		  <label for="exampleInputEmail1">Title</label>
		  <input name="title" class="form-control" placeholder="Post Title">
		  <small class="form-text text-muted">Make it catchy!</small>
		</div>
		<input style="display: none" id="redirectTo" name="redirectTo">
		<script>
			// This does not seem to work yet
			document.getElementById("redirectTo").value = window.location.href;
			console.log(document.getElementById("redirectTo"))
		</script>
		<div class="form-group">
		  <label for="exampleInputPassword1">Content</label>
		  <textarea name="content" class="form-control" placeholder="Post Content"></textarea>
		  <small class="form-text text-muted">You can use html fragments in the post content. Be careful!</small>
		</div>
		<div class="form-group">
			<label for="exampleInputPassword1">Tags</label>
			<input name="tagNames" class="form-control" placeholder="Tags">
			<small class="form-text text-muted">Comma seaprated list of tags?</small>
	  	</div>
		<div class="form-group">
		  <label for="exampleInputPassword1">Password</label>
		  <input name="password" class="form-control" placeholder="Password">
		  <small class="form-text text-muted">Password for posting content. Did you think you can just post away without it?</small>
		</div>
		<button type="submit" class="btn btn-primary">Submit</button>
  		</form>
	</div>
`
var EditPostBody = `
	<h1>
		Edit Post
	</h1>
	<div id="content" style="text-align: left;">
	<form action="api/post" method="post">
		<input name="id" style="display: none" value="{{ .Post.Id }}" class="form-control">
		<div class="form-group">
		  <label for="exampleInputEmail1">Title</label>
		  <input name="title" class="form-control" value="{{ .Post.Title }}" placeholder="Post Title">
		  <small class="form-text text-muted">Make it catchy!</small>
		</div>
		<input style="display: none" id="redirectTo" name="redirectTo">
		<script>
			// This does not seem to work yet
			document.getElementById("redirectTo").value = window.location.href;
			console.log(document.getElementById("redirectTo"))
		</script>
		<div class="form-group">
		  <label for="exampleInputPassword1">Content</label>
		  <textarea name="content" class="form-control" placeholder="Post Content">{{ .Post.Content }}</textarea>
		  <small class="form-text text-muted">You can use html fragments in the post content. Be careful!</small>
		</div>
		<div class="form-group">
			<label for="exampleInputPassword1">Tags</label>
			<input name="tagNames" class="form-control" value="{{ .TagNames }}" placeholder="Tags">
			<small class="form-text text-muted">Comma seaprated list of tags?</small>
	  	</div>
		<div class="form-group">
		  <label for="exampleInputPassword1">Password</label>
		  <input name="password" class="form-control" placeholder="Password">
		  <small class="form-text text-muted">Password for posting content. Did you think you can just post away without it?</small>
		</div>
		<button type="submit" class="btn btn-primary">Submit</button>
  		</form>
	</div>
`
