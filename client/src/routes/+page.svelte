<script>
    import { deserialize } from '$app/forms';
	import { goto } from '$app/navigation';
	
	let username = $state('jioo');
	let password = $state('password');

	const handleLogin = async (e) => {
        e.preventDefault();
        
		try {
            const formData = new FormData();
            formData.append('username', username);
            formData.append('password', password);
			const response = await fetch(`?/login`, {
				method: 'POST',
				body: formData
			});
			const result = deserialize(await response.text());
            if (result.type === 'success') {
                goto('/albums');
            }

		} catch (error) {
			console.error(error);
			alert(error.message)
		}
	};
</script>

<div class="hero min-h-screen">
	<div class="hero-content flex-col lg:flex-row-reverse">
		<div class="card bg-base-300 w-full max-w-sm flex-shrink-0 shadow-2xl">
			<div class="card-body">
				<div class="form-control">
					<label class="label" for="username">
						<span class="label-text">Username</span>
					</label>
					<input
						type="text"
						id="username"
						bind:value={username}
						class="input input-bordered"
						required
					/>
				</div>
				<div class="form-control">
					<label class="label" for="password">
						<span class="label-text">Password</span>
					</label>
					<input
						type="password"
						id="password"
						bind:value={password}
						class="input input-bordered"
						required
					/>

                    <div class="mt-1">
                        <a href="/register" class="label-text-alt link link-hover">Don't have an account?</a>
                    </div>
				</div>
				<div class="form-control mt-4">
					<button class="btn btn-primary w-full" onclick={handleLogin}>Login</button>
				</div>
			</div>
		</div>
	</div>
</div>
