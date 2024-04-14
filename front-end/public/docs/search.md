**Telemetry data search block/form**

This block provides a form to find all telemetry data received from the device.

**Process of filling out the form:**
- Select time periods (start and end).
- Select telemetry type.
- Select the device (after selection, its description will be displayed).
- Select the type of information display:
    - **Map:** Shows data only if you select geolocation as the telemetry type. Clicking on a point on the map will give you a more detailed description of the other selected telemetry.
    - **Graph:** Shows all available data that you have selected. It is possible to filter certain telemetry.
- Consider ignored data (checkbox): Will also display ignored data when parsing on the server side. If you select dates when the device was just turned on, it may transmit incomplete or incorrect data, which will be ignored to prevent incorrect information being displayed.

**Button functionality:**
- **Clear:** Returns the form to its original empty state.
- **Save Query:** Saves the current form to history (see history help).
- **Find:** Performs a search.