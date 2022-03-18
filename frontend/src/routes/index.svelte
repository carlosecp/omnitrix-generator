<script context="module">
	export const load = async ({ fetch }) => {
		const res = await fetch("http://127.0.0.1:8080");
		const aliens = await res.json();

		return {
			props: { aliens }
		}
	}
</script>

<script>
	export let aliens;

	import AlienCard from "../lib/AlienCard.svelte";

	let searchTerm = "";

	$: filteredAliens = aliens.filter(alien => {
		return alien.name.toLowerCase().includes(searchTerm) ||
			alien.species.toLowerCase().includes(searchTerm) ||
			alien.homePlanet.toLowerCase().includes(searchTerm);
	})
</script>

<div class="p-4 space-y-8">
	<img src="/img/ben10.png" class="mx-auto" />
	<label class="block mx-auto max-w-lg">
		<input
			bind:value={searchTerm}
			class="w-full px-3 py-2 bg-white border border-slate-300 rounded-md focus:ring text-sm shadow-sm placeholder-slate-400 focus:outline-none focus:border-accent focus:ring-accent"
			placeholder="Search an alien by name, species or planet..."/>
	</label>
</div>
<main class="mt-4 p-4 grid sm:grid-cols-2 gap-4">
	{#each filteredAliens as alien}
		<AlienCard {alien} />
	{/each}
</main>
