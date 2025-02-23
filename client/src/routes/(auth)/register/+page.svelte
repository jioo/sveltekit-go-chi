<script>
    import { deserialize } from '$app/forms';
	import { goto } from '$app/navigation';
	
	let user = $state({
		firstName: '',
		lastName: '',
		username: '',
		password: ''
	});
	let errors = $state([]);
	
	const handleRegister = async (e) => {
        e.preventDefault();
        
		try {
            const formData = new FormData();
            formData.append('firstName', user.firstName);
            formData.append('lastName', user.lastName);
            formData.append('username', user.username);
            formData.append('password', user.password);
			const response = await fetch(`?/register`, {
				method: 'POST',
				body: formData
			});
			
			const result = deserialize(await response.text());
            const { data } = result;
            if (data.errors) {
                errors = data.errors;
                return false;
            }

			goto('/login');

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
                {#if errors.length}
                    <div role="alert" class="alert alert-error">
                        <ul class="list-disc list-inside space-y-1">
                            {#each errors as error}
                                <li class="text-sm">{error.message}</li>
                            {/each}
                        </ul>
                    </div>
                {/if}

				<div class="form-control">
					<label class="label" for="first-name">
						<span class="label-text">First Name *</span>
					</label>
					<input
						type="text"
						id="first-name"
						bind:value={user.firstName}
						class="input input-bordered"
					/>
				</div>

				<div class="form-control">
					<label class="label" for="last-name">
						<span class="label-text">Last Name *</span>
					</label>
					<input
						type="text"
						id="last-name"
						bind:value={user.lastName}
						class="input input-bordered"
					/>
				</div>

				<div class="form-control">
					<label class="label" for="username">
						<span class="label-text">Username *</span>
					</label>
					<input
						type="text"
						id="username"
						bind:value={user.username}
						class="input input-bordered"
					/>
				</div>
				<div class="form-control">
					<label class="label" for="password">
						<span class="label-text">Password *</span>
					</label>
					<input
						type="password"
						id="password"
						bind:value={user.password}
						class="input input-bordered"
					/>
				</div>
				<div class="form-control mt-4">
					<button class="btn btn-primary w-full mb-4" onclick={handleRegister}>Register</button>
					<a href="/login" class="btn btn-secondary w-full">Cancel</a>
				</div>
			</div>
		</div>
	</div>
</div>
