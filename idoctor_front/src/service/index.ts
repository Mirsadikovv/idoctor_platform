export * from "@module/Language/service";

export * from "@module/TranslatedContent/service";

export * from "@module/User/service";

export * from "@module/Role/service";

export * from "@module/Auth/service";

export * from "@module/Device/service";

export * from "@module/Problem/service";

export * from "@module/Supplier/service";

export * from "@module/Part/service";

export * from "@module/Order/service";

export * from "@module/Media/service/media.service";

export type PageDataType<T> = {
	totalRows: number;
	totalPages: number;
	pageSize: number;
	currentPage: number;
	data: T[];
};

export type IdType = {
	id: number | string;
};
