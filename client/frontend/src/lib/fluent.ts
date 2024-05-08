import {
	provideFluentDesignSystem,
	baseLayerLuminance,
	fluentButton, fluentTextField
} from "@fluentui/web-components"

baseLayerLuminance.setValueFor(document.body, 0)

provideFluentDesignSystem().register(fluentButton(), fluentTextField())
