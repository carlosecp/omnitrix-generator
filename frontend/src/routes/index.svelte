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

<div class="m-4 grid sm:grid-cols-2 gap-4">
	{#each aliens as alien}
		<article class="p-4 border rounded">
			<div class="flex gap-4">
				<img src={alien.imgURL} class="h-44 w-40 object-cover object-top" />
				<div class="text-neutral-600">
					<h1 class="font-medium text-xl text-neutral-900">{alien.name}</h1>
					<strong class="font-medium text-neutral-800">Species:</strong> {alien.species}
					<br/>
					<strong class="font-medium text-neutral-800">Home Planet:</strong> {alien.homePlanet}
					<br/>
					<strong class="font-medium text-neutral-800">Powers:</strong> {alien.homePlanet}
					<ul>
						{#each alien.powers as power}
							<li>{power}</li>
						{/each}
					</ul>
				</div>
			</div>
		</article>
	{/each}
</div>
