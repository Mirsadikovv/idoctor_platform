<script setup lang="ts">
import { computed, ref } from "vue";
import Title from "./Title.vue";
import { getStatusColor, formatDateToUtc } from "@/modules/Appeal/shared/utils";
import {
	EWorkflowAction,
	WorkflowService,
	type TWorkflowDetail,
} from "@/modules/Appeal/shared/service";
import { dateFormat } from "@/common";

const workflowData = ref<TWorkflowDetail | null>(null);
const loading = ref(false);
const modelValue = ref(false);
const appealId = ref(0);
const workflowId = ref(0);

// Загрузка данных при открытии модалки
async function fetchWorkflowData() {
	if (!modelValue.value) return;

	loading.value = true;
	const data = await WorkflowService.getById(appealId.value, workflowId.value);
	workflowData.value = data;
	loading.value = false;
}

function openDialog(_appealId: number, _workflowId: number) {
	modelValue.value = true;
	appealId.value = _appealId;
	workflowId.value = _workflowId;
	fetchWorkflowData();
}

function closeDialog() {
	modelValue.value = false;
}

const signer = computed(() => {
	return workflowData.value?.Signer?.signer;
});

const appeal = computed(() => {
	return workflowData.value?.SignData?.body?.appeal;
});
const workflowAction = computed(() => {
	return workflowData.value?.SignData?.body?.workflow_action;
});

const currentWorkflow = computed(() => {
	return workflowData.value;
});

// Получение цвета для статуса

defineExpose({
	openDialog,
	closeDialog,
});
</script>

<template>
	<q-dialog
		v-model="modelValue"
		persistent
		maximized
		transition-show="slide-up"
		transition-hide="slide-down"
	>
		<q-card class="bg-grey-1">
			<q-card-section class="row items-center bg-primary text-white sticky top-0 z-10">
				<Title class="text-white mb-0">
					<q-icon name="verified_user" class="mr-2" />
					{{ $tl("signature_information") }}
					<q-space />
					<q-btn icon="close" flat round dense @click="modelValue = false" />
				</Title>
			</q-card-section>

			<!-- Loader во время загрузки -->
			<q-card-section
				v-if="loading"
				class="q-pa-xl flex justify-center items-center"
				style="min-height: 400px"
			>
				<q-spinner color="primary" size="3em" />
			</q-card-section>

			<q-card-section v-else-if="!workflowData" class="q-pa-md text-center text-grey-6">
				{{ $tl("no_data") }}
			</q-card-section>

			<q-card-section v-else class="q-pa-md">
				<div class="row q-col-gutter-md">
					<!-- Основная информация о workflow -->
					<div class="col-12">
						<q-card flat bordered>
							<q-card-section class="bg-purple-1">
								<div class="text-h6 text-purple-9 flex items-center">
									<q-icon name="timeline" size="sm" class="mr-2" />
									{{ $tl("workflow_information") }}
								</div>
							</q-card-section>
							<q-separator />
							<q-card-section v-if="currentWorkflow" class="q-pa-none">
								<q-markup-table separator="horizontal" flat bordered>
									<tbody>
										<tr>
											<td class="font-bold text-grey-7" style="width: 30%">
												{{ $tl("appeal_id") }}
											</td>
											<td>
												{{ currentWorkflow.appealId?.toString() || "-" }}
											</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("created_at") }}
											</td>
											<td>
												{{
													dateFormat(
														currentWorkflow.createdAt,
														"DD.MM.YYYY HH:mm:ss",
													)
												}}
											</td>
										</tr>
										<tr
											v-if="
												workflowData.action === EWorkflowAction.REJECT ||
												workflowData.action === EWorkflowAction.COMPLETED
											"
										>
											<td class="font-bold text-grey-7">
												{{ $tl("current_inspector") }}
											</td>
											<td>
												{{ currentWorkflow.currentInspectorName || "-" }}
											</td>
										</tr>
										<tr
											v-if="
												workflowData.action === EWorkflowAction.REJECT ||
												workflowData.action === EWorkflowAction.COMPLETED
											"
										>
											<td class="font-bold text-grey-7">
												{{ $tl("current_organization") }}
											</td>
											<td>{{ currentWorkflow.currentOrgName || "-" }}</td>
										</tr>
										<tr v-if="workflowData.action === EWorkflowAction.APPROVED">
											<td class="font-bold text-grey-7">
												{{ $tl("main_inspector_name") }}
											</td>
											<td>{{ currentWorkflow.fromInspectorName || "-" }}</td>
										</tr>
										<tr v-if="workflowData.action === EWorkflowAction.APPROVED">
											<td class="font-bold text-grey-7">
												{{ $tl("organization_name") }}
											</td>
											<td>{{ currentWorkflow.fromOrgName || "-" }}</td>
										</tr>
										<tr v-if="workflowData.action === EWorkflowAction.APPROVED">
											<td class="font-bold text-grey-7">
												{{ $tl("to_inspector_name") }}
											</td>
											<td>{{ currentWorkflow.toInspectorName || "-" }}</td>
										</tr>
										<tr v-if="currentWorkflow.comment">
											<td class="font-bold text-grey-7">
												{{ $tl("comment") }}
											</td>
											<td style="white-space: pre-wrap">
												{{ currentWorkflow.comment }}
											</td>
										</tr>
										<tr v-if="currentWorkflow.systemComment">
											<td class="font-bold text-grey-7">
												{{ $tl("system_comment") }}
											</td>
											<td style="white-space: pre-wrap">
												{{ currentWorkflow.systemComment }}
											</td>
										</tr>
									</tbody>
								</q-markup-table>
							</q-card-section>
						</q-card>
					</div>

					<!-- Информация о подписанте -->
					<div class="col-12">
						<q-card flat bordered>
							<q-card-section class="bg-blue-1">
								<div class="text-h6 text-blue-9 flex items-center">
									<q-icon name="person" size="sm" class="mr-2" />
									{{ $tl("signer_information") }}
								</div>
							</q-card-section>
							<q-separator />
							<q-card-section v-if="signer" class="q-pa-none">
								<q-markup-table separator="horizontal" flat bordered>
									<tbody>
										<tr>
											<td class="font-bold text-grey-7" style="width: 30%">
												{{ $tl("full_name") }}
											</td>
											<td>{{ signer.Name || "-" }}</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">{{ $tl("tin") }}</td>
											<td>{{ signer.Tin || "-" }}</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("certificate_number") }}
											</td>
											<td>{{ signer.No?.toString() || "-" }}</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("ip_address") }}
											</td>
											<td>{{ signer.Ip || "-" }}</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("signature_time") }}
											</td>
											<td>{{ formatDateToUtc(signer.Time) || "-" }}</td>
										</tr>
									</tbody>
								</q-markup-table>
							</q-card-section>
							<q-card-section v-else>
								<div class="text-center text-grey-6">
									{{ $tl("no_signer_data") }}
								</div>
							</q-card-section>
						</q-card>
					</div>

					<!-- Информация о workflow действии -->

					<!-- Краткая информация об обращении -->
					<div class="col-12" v-if="appeal">
						<q-card flat bordered>
							<q-card-section class="bg-orange-1">
								<div class="text-h6 text-orange-9 flex items-center">
									<q-icon name="description" size="sm" class="mr-2" />
									{{ $tl("appeal_information") }}
								</div>
							</q-card-section>
							<q-separator />
							<q-card-section class="q-pa-none">
								<q-markup-table separator="horizontal" flat bordered>
									<tbody>
										<tr>
											<td class="font-bold text-grey-7" style="width: 30%">
												{{ $tl("old_status") }}
											</td>
											<td>
												<q-chip
													:color="getStatusColor(appeal.status)"
													text-color="white"
													icon="info"
												>
													{{ $tl(appeal.status || "") }}
												</q-chip>
											</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("new_status") }}
											</td>
											<td>
												<q-chip
													:color="getStatusColor(workflowAction?.action)"
													text-color="white"
													icon="info"
												>
													{{ $tl(workflowAction?.action || "") }}
												</q-chip>
											</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("source") }}
											</td>
											<td>
												<q-chip
													color="indigo"
													text-color="white"
													icon="source"
												>
													{{ $tl(appeal.source || "") }}
												</q-chip>
											</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("full_name") }}
											</td>
											<td>
												{{ appeal.lastName }} {{ appeal.firstName }}
												{{ appeal.middleName }}
											</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("phone_number") }}
											</td>
											<td>{{ appeal.phoneNumber || "-" }}</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("company_name") }}
											</td>
											<td>{{ appeal.companyName || "-" }}</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("company_inn") }}
											</td>
											<td>{{ appeal.companyINN || "-" }}</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("inspector_name") }}
											</td>
											<td>{{ appeal.inspectorName || "-" }}</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("created_at") }}
											</td>
											<td>
												{{
													dateFormat(
														appeal.createdAt,
														"DD.MM.YYYY HH:mm:ss",
													)
												}}
											</td>
										</tr>
										<tr v-if="appeal.industry">
											<td class="font-bold text-grey-7">
												{{ $tl("industry") }}
											</td>
											<td>{{ appeal.industry.name || "-" }}</td>
										</tr>
										<tr
											v-if="
												appeal.complaintTypes &&
												appeal.complaintTypes.length > 0
											"
										>
											<td class="font-bold text-grey-7">
												{{ $tl("complaint_types") }}
											</td>
											<td>
												<div class="flex gap-2 flex-wrap q-py-sm">
													<q-chip
														v-for="type in appeal.complaintTypes"
														:key="type.id"
														color="teal"
														text-color="white"
													>
														{{ type.name }}
													</q-chip>
												</div>
											</td>
										</tr>
										<tr>
											<td class="font-bold text-grey-7">
												{{ $tl("complaint_text") }}
											</td>
											<td style="white-space: pre-wrap">
												{{ appeal.complaintTxt }}
											</td>
										</tr>
									</tbody>
								</q-markup-table>
							</q-card-section>
						</q-card>
					</div>
				</div>
			</q-card-section>

			<q-card-actions align="right" class="bg-white sticky bottom-0 q-pa-md border-t">
				<q-btn
					flat
					:label="$tl('close')"
					color="primary"
					@click="modelValue = false"
					icon="close"
					size="md"
					class="px-6"
				/>
			</q-card-actions>
		</q-card>
	</q-dialog>
</template>

<style scoped lang="scss">
.border-t {
	border-top: 1px solid rgba(0, 0, 0, 0.12);
}
</style>
