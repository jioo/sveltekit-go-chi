<script>
    import { deserialize } from '$app/forms';

    const { data } = $props();
    const album = $state(data);

    const handleUpdate = async (e) => {
        e.preventDefault();
        
		try {
            const formData = new FormData();
            formData.append('title', album.title);
            formData.append('artist', album.artist);
            formData.append('price', album.price);
			const response = await fetch(`?/save`, {
				method: 'POST',
				body: formData
			});
			const result = deserialize(await response.text());
            console.log(result)

		} catch (error) {
			console.error(error);
		}
	};
</script>

<div class="bg-base-300 p-6">
    <h1 class="text-2xl font-bold">Edit Album</h1>
    <div class="breadcrumbs mb-4 text-sm">
        <ul>
            <li><a class="link text-xs" href="/albums">Albums</a></li>
        </ul>
    </div>
    
    <div class="grid" >
        <fieldset class="fieldset">
            <legend class="fieldset-legend">Title *</legend>
            <input
                class="input"
                type="text"
                name="title"
                bind:value={album.title}
            />
        </fieldset>
    
        <fieldset class="fieldset">
            <legend class="fieldset-legend">Artist *</legend>
            <input
                class="input"
                type="text"
                name="artist"
                bind:value={album.artist}
            />
        </fieldset>
    
        <fieldset class="fieldset">
            <legend class="fieldset-legend">Price *</legend>
            <input
                class="input text-right"
                type="text"
                name="price"
                bind:value={album.price}
            />
        </fieldset>

        <button class="btn btn-primary w-80 mt-4" onclick={handleUpdate}>Update</button>
    </div>
</div>