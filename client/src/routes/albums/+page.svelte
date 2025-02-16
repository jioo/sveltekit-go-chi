<script>
    import { deserialize } from '$app/forms';

    const { data } = $props();
    let albums = $state(data.albums);
    let modalDelete;
    let pending;

    const confirmDelete = async (e) => {
        e.preventDefault();

        try {
            const formData = new FormData();
            formData.append('id', pending.id);
            const response = await fetch(`?/delete`, {
                method: 'POST',
                body: formData
            });
            const result = deserialize(await response.text());
            if (result.type === 'success') {
                albums = albums.filter((album) => album.id !== pending.id);
            }
            modalDelete.close();

        } catch (error) {
            console.error(error);
        }
    };
</script>

<div class="bg-base-300 p-6">

    <h1 class="text-2xl font-bold mb-4">Albums</h1>

    <a href="albums/0" class="btn btn-primary mb-4">Create</a>

    <table class="table table-lg bg-base-200">
        <thead>
            <tr>
                <th>ID</th>
                <th>Title</th>
                <th>Artist</th>
                <th>Price</th>
                <th>Action</th>
            </tr>
        </thead>
        <tbody>
            {#each albums as album}
                <tr>
                    <td>{album.id}</td>
                    <td>{album.title}</td>
                    <td>{album.artist}</td>
                    <td>{album.price}</td>
                    <td>
                        <a href="albums/{album.id}" class="btn btn-sm btn-primary">Edit</a>
                        <button 
                            class="btn btn-sm btn-error" 
                            onclick={() => {
                                pending = { id: album.id };
                                modalDelete.showModal();
                            }}
                        >
                            Delete
                        </button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>

<dialog bind:this={modalDelete} class="modal">
	<div class="modal-box">
		<p>Are you sure you want to delete this?</p>
		<div class="modal-action">
			<button class="btn" onclick={() => modalDelete.close()}>Cancel</button>
			<button class="btn btn-error" onclick={confirmDelete}>Delete</button>
		</div>
	</div>
</dialog>