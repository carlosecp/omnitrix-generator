<script>
	export let alien;
</script>

<article class="p-4 flex flex-col sm:flex-row items-center sm:items-stretch gap-4 bg-white border border-slate-300 rounded-md shadow-sm">
	<img src={alien.imgURL} class="h-44 sm:w-40 object-cover object-top" />
	<div class="alien-card-content">
		<h1 class="mb-2 font-semibold text-xl text-gray-700">{alien.name}</h1>
		<div><strong>Species:</strong> {alien.species}</div>
		<div><strong>Home Planet:</strong> {alien.homePlanet}</div>
		<div><strong>Powers:</strong> {alien.homePlanet}</div>
	</div>
</article>
