// Lazy imports будут загружаться только при вызове функции

interface ExportToExcelOptions {
	data: any[];
	visibleColumns: string[];
	columnMap: Record<string, string>;
	filename?: string;
	excludedColumns?: string[];
}

export async function exportToExcel({
	data,
	visibleColumns,
	columnMap,
	filename = "export.xlsx",
	excludedColumns = [],
}: ExportToExcelOptions) {
	try {
		console.log("ExportToExcel called with:", { data, visibleColumns, columnMap, filename, excludedColumns });
		
		// Lazy load ExcelJS только при необходимости
		const ExcelJS = (await import("exceljs")).default;
		const { saveAs } = await import("file-saver");
		
		const workbook = new ExcelJS.Workbook();
		const worksheet = workbook.addWorksheet("Statistics");

		// === Header row ===
		const headers = visibleColumns
			.filter((col) => !excludedColumns.includes(col))
			.map((col) => ({
				header: columnMap[col] || col,
				key: col,
				width: 25,
			}));

		console.log("Generated headers:", headers);
		worksheet.columns = headers;

		// === Data rows ===
		console.log("Processing data rows:", data.length);
		for (const row of data) {
			const rowData: Record<string, any> = {};
			for (const col of visibleColumns) {
				if (excludedColumns.includes(col)) continue;
				rowData[col] =
					row[col] ??
					row[col.toLowerCase()] ??
					row[col.replace(/([A-Z])/g, "_$1").toLowerCase()] ??
					"";
			}
			worksheet.addRow(rowData);
		}
		console.log("Added", data.length, "rows to worksheet");

		// === Header styling ===
		worksheet.getRow(1).eachCell((cell) => {
			cell.font = { bold: true, color: { argb: "FFFFFFFF" } };
			cell.fill = {
				type: "pattern",
				pattern: "solid",
				fgColor: { argb: "4472C4" },
			};
			cell.alignment = { horizontal: "center" };
		});

		// === Создание и загрузка файла ===
		const buffer = await workbook.xlsx.writeBuffer();
		const blob = new Blob([buffer], {
			type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		});
		saveAs(blob, filename);
	} catch (err) {
		console.error("Excel eksportida xatolik:", err);
	}
}
