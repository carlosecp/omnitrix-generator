<script context="module">
	export const load = async ({ fetch }) => {
		const res = await fetch("http://127.0.0.0:8080");
		const aliens = await res.json();

		return {
			props: { aliens }
		}
	}
</script>

<script>
	export let aliens;
</script>

<div class="grid">
	{#each aliens as alien}
		<div>{alien.Name}</div>
		<img src={alien.ImgURL} />
	{/each}
</div>
