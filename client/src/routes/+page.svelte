<script>
    const { data } = $props();
    const albums = $state(data.albums);
</script>

<div class="bg-base-300 p-6">

    <h1 class="text-2xl font-bold mb-4">Albums</h1>

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
                        <a href="/{album.id}" class="btn btn-sm btn-primary">Edit</a>
                        <button class="btn btn-sm btn-error">Delete</button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>