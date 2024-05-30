<script>
	export let name;
	import { onMount } from "svelte";
	let data = null;

	onMount(async () => {
		// Read json from songs.json
		const response = await fetch("/songs.json");
		// Parse data as json
		data = await response.json();
		console.log(data);
	});
</script>

<main>
	<h1>Hello {name}!</h1>
	<p>
		Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn
		how to build Svelte apps.
	</p>
</main>

{#if data}
	<div
		class="container"
		style="width: 100%; display: flex; flex-direction: column; align-items: center;"
	>
		{#each data as item}
			<div class="song" style="text-align: center;">
				<h2>{item.title}</h2>
				<p>{item.artist}</p>
				<p>{item.link}</p>
				<p>{item.date}</p>
			</div>
		{/each}
	</div>
{/if}

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
