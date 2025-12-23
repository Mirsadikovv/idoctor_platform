import { registerRoutes } from "@/router/layout";
import { UserRoutes } from "./User/routes";
import { emptyRoute } from "./Auth/router";
import { forbiddenPage, notfoundPage } from "./Error/router/error";
import { LanguageGroupRoutes } from "./languages";
import { DeviceRoutes } from "./Device/routes";
import { ProblemRoutes } from "./Problem/routes";
import { SupplierRoutes } from "./Supplier/routes";
import { PartRoutes } from "./Part/routes";
import { RoleRoutes } from "./Role/routes";
import { OrderRoutes } from "./Order/routes";

registerRoutes(
	emptyRoute,

	// Error
	notfoundPage,
	forbiddenPage,

	// User sort 1
	...UserRoutes(1),

	// Role sort 2
	...RoleRoutes(2),

	// Device sort 3
	...DeviceRoutes(3),

	// Problem sort 4
	...ProblemRoutes(4),

	// Supplier sort 5
	...SupplierRoutes(5),

	// Part sort 6
	...PartRoutes(6),

	// Order sort 7
	...OrderRoutes(7),

	// Languages sort 10
	LanguageGroupRoutes(10),
);
