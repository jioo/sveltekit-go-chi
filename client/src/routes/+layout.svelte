<script>
	import '../app.css';
	let { children } = $props();
</script>

<div class="flex min-h-screen flex-col justify-between">
	<main class="container mx-auto mb-auto mt-2 px-4">
		{@render children()}
	</main>
</div>