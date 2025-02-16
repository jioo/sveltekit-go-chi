<script>
    import { enhance } from '$app/forms';
    const { form, data } = $props();
    const album = $state(data);
</script>

<div class="bg-base-300 p-6">
    <h1 class="text-2xl font-bold">Edit Album</h1>
    <div class="breadcrumbs mb-4 text-sm">
        <ul>
            <li><a class="link text-xs" href="/">Albums</a></li>
        </ul>
    </div>
    
    <form use:enhance method="POST" action="?/save" class="grid" >
        <fieldset class="fieldset">
            <legend class="fieldset-legend">Title *</legend>
            <input
                class="input"
                type="text"
                name="title"
                value={album.title}
            />
        </fieldset>
    
        <fieldset class="fieldset">
            <legend class="fieldset-legend">Artist *</legend>
            <input
                class="input"
                type="text"
                name="artist"
                value={album.artist}
            />
        </fieldset>
    
        <fieldset class="fieldset">
            <legend class="fieldset-legend">Price *</legend>
            <input
                class="input text-right"
                type="text"
                name="price"
                value={album.price}
            />
        </fieldset>

        <button class="btn btn-primary w-80 mt-4" type="submit">Update</button>
    </form>
</div>