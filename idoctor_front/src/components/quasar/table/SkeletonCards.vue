<template>
	<div class="mobile-skeleton-container">
		<q-card 
			v-for="n in 6" 
			:key="n" 
			class="mobile-skeleton-card q-mb-md"
			flat
			bordered
		>
			<q-card-section>
				<div class="row q-col-gutter-md">
					<div class="col-12">
						<q-skeleton type="text" width="60%" />
					</div>
					<div class="col-6">
						<q-skeleton type="text" width="80%" />
					</div>
					<div class="col-6">
						<q-skeleton type="text" width="70%" />
					</div>
					<div class="col-12">
						<q-skeleton type="text" width="90%" />
					</div>
				</div>
			</q-card-section>
		</q-card>
	</div>
</template>

<style scoped lang="scss">
.mobile-skeleton-container {
	padding: 16px 0;
}

.mobile-skeleton-card {
	border-radius: 8px;
	box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}
</style>