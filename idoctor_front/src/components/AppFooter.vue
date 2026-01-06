<script setup lang="ts">
import { ref } from "vue";
import ButtonDialog from "@/components/quasar/dialog/ButtonDialog.vue";
import type { LanguageType } from "@/service";

interface Props {
	username?: string;
	languages?: LanguageType[];
	currentLanguageId?: number;
	showAddButton?: boolean;
	addButtonRoute?: any;
	addButtonIcon?: string;
	showBackButton?: boolean;
	backButtonRoute?: object;
	centerActions?: boolean;
}

withDefaults(defineProps<Props>(), {
	showAddButton: false,
	addButtonIcon: "add_circle",
	showBackButton: false,
	centerActions: true,
});

const emit = defineEmits<{
	toggleDrawer: [];
	goToProfile: [];
	setLang: [language: LanguageType];
	showLogoutConfirm: [show: boolean];
	logout: [];
}>();

const mobileLogoutConfirm = ref(false);
</script>

<template>
	<q-footer class="clean-header shadow-lg border-b border-white/20">
		<q-toolbar class="h-18 gap-x-3 clean-toolbar bg-transparent text-white px-6">
			<q-btn
				size="lg"
				icon="menu"
				@click="$emit('toggleDrawer')"
				class="clean-btn bg-transparent! text-white! rounded-lg! transition-all duration-200 min-h-44px! min-w-44px! hover:bg-white/15! active:bg-white/20!"
				dense
				round
				flat
			/>
			<div class="flex-1" :class="centerActions ? 'flex justify-center' : 'flex gap-2'">
				<q-btn
					v-if="showBackButton"
					size="lg"
					icon="arrow_back"
					:to="backButtonRoute"
					class="clean-btn bg-transparent! text-white! rounded-lg! transition-all duration-200 min-h-44px! min-w-44px! hover:bg-white/15! active:bg-white/20!"
					dense
					round
					flat
				/>
				<q-btn
					v-if="showAddButton || $canPage(addButtonRoute?.name || '')"
					size="lg"
					:icon="addButtonIcon"
					:to="addButtonRoute"
					class="clean-btn bg-transparent! text-white! rounded-lg! transition-all duration-200 min-h-44px! min-w-44px! hover:bg-white/15! active:bg-white/20!"
					dense
					round
					flat
				/>
				<slot name="actions"></slot>
			</div>

			<!-- Mobile Menu -->
			<q-btn-dropdown
				flat
				aria-label="Menu"
				size="lg"
				dropdown-icon="settings"
				class="flex clean-menu-dropdown bg-transparent! text-white! rounded-lg! transition-all duration-200 min-h-44px! min-w-44px!"
			>
				<q-list
					class="telegram-menu-list min-w-200px max-w-320px bg-white/95 backdrop-blur-20 rounded-xl border border-black/10 shadow-2xl p-2"
				>
					<!-- Profile Item -->
					<q-item
						clickable
						@click="$emit('goToProfile')"
						class="telegram-menu-item rounded-lg my-1 transition-all duration-200 min-h-48px bg-green-50 border border-green-100 hover:bg-green-100! active:bg-green-150!"
						v-close-popup
					>
						<q-item-section avatar>
							<q-icon name="account_box" color="positive" size="20px" />
						</q-item-section>
						<q-item-section>
							<q-item-label class="text-sm font-medium text-gray-800">
								{{ username }}
							</q-item-label>
						</q-item-section>
					</q-item>

					<q-separator class="my-1 bg-gray-200" />

					<!-- Language Selector -->
					<q-expansion-item
						icon="language"
						:label="$tl('language')"
						header-class="text-gray-800! font-medium px-3 py-2 rounded-lg hover:bg-gray-50!"
						class="telegram-menu-item my-1"
					>
						<q-item
							v-for="language of languages"
							:key="language.id"
							clickable
							:class="[
								'mx-2 my-1 rounded-lg min-h-40px transition-colors',
								language.id === currentLanguageId
									? 'bg-blue-100! text-blue-800! font-semibold'
									: 'hover:bg-blue-50!',
							]"
							@click="$emit('setLang', language)"
							v-close-popup
						>
							<q-item-section>
								<q-item-label class="text-sm">
									{{ language.name }}
								</q-item-label>
							</q-item-section>
							<q-item-section side>
								<q-icon
									v-if="language.id === currentLanguageId"
									name="check_circle"
									size="16px"
									color="positive"
								/>
							</q-item-section>
						</q-item>
					</q-expansion-item>

					<q-separator class="my-1 bg-gray-200" />

					<!-- Logout Mobile -->
					<q-item
						clickable
						@click="mobileLogoutConfirm = true"
						class="telegram-menu-item rounded-lg my-1 transition-all duration-200 min-h-48px bg-red-50 border border-red-100 hover:bg-red-100! active:bg-red-150!"
						v-close-popup
					>
						<q-item-section avatar>
							<q-icon name="logout" color="negative" size="20px" />
						</q-item-section>
						<q-item-section>
							<q-item-label class="text-sm font-medium text-red-600">
								{{ $tl("logout") }}
							</q-item-label>
						</q-item-section>
					</q-item>
				</q-list>
			</q-btn-dropdown>

			<!-- Logout desktop -->
			<ButtonDialog
				:classBtn="'hidden lg:flex! md:hidden! bg-red-600 text-white rounded-lg px-3 py-2 transition-all duration-200 hover:bg-red-700! hover:-translate-y-0.5 hover:shadow-lg active:translate-y-0'"
				icon="logout"
				label="logout"
			>
				<q-card class="min-w-300px w-400px rounded-xl overflow-hidden">
					<q-card-section
						class="bg-gradient-to-r from-blue-600 to-blue-700 text-white p-6"
					>
						<div class="text-2xl font-bold">{{ $tl("logout_confirm") }}</div>
						<div class="text-lg opacity-90 mt-1">{{ $tl("are_you_sure") }}?</div>
					</q-card-section>

					<q-card-actions align="center" class="flex gap-3 p-4">
						<q-btn
							no-caps
							outline
							class="flex-1 py-2 px-4 border-2 border-blue-500 text-blue-600 rounded-lg font-medium hover:bg-blue-50 transition-colors"
							:label="$tl('no')"
							v-close-popup
						/>
						<q-btn
							no-caps
							class="flex-1 py-2 px-4 bg-red-600 text-white rounded-lg font-medium hover:bg-red-700 transition-colors"
							:label="$tl('yes')"
							v-close-popup
							@click="$emit('logout')"
						/>
					</q-card-actions>
				</q-card>
			</ButtonDialog>
		</q-toolbar>

		<!-- Mobile Logout Confirmation Dialog -->
		<q-dialog v-model="mobileLogoutConfirm" persistent>
			<q-card class="min-w-300px w-400px rounded-xl overflow-hidden">
				<q-card-section class="bg-gradient-to-r from-blue-600 to-blue-700 text-white p-6">
					<div class="text-2xl font-bold">{{ $tl("logout_confirm") }}</div>
					<div class="text-lg opacity-90 mt-1">{{ $tl("are_you_sure") }}?</div>
				</q-card-section>

				<q-card-actions align="center" class="flex gap-3 p-4">
					<q-btn
						no-caps
						outline
						color="secondary"
						class="flex-1 py-2 px-4"
						:label="$tl('no')"
						v-close-popup
					/>
					<q-btn
						no-caps
						outline
						color="negative"
						class="flex-1 py-2 px-4"
						:label="$tl('yes')"
						v-close-popup
						@click="emit('logout')"
					/>
				</q-card-actions>
			</q-card>
		</q-dialog>
	</q-footer>
</template>
