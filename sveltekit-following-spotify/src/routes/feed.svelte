<script>
    import { onMount } from "svelte";
    let data = [];

    onMount(async () => {
        const response = await fetch("./feed.xml.js");
        print(response);
        data = await response.json();
    });
</script>

{#each data as item (item.id)}
    <div class="box">
        <h2>{item.title}</h2>
        <p>{item.description}</p>
    </div>
{/each}
