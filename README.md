# urban-computing-machine
API que convierte unidades de medida

## Épica: 
Conversión de Unidades

## User Stories:

1. Como usuario, quiero poder convertir unidades de longitud (metros a pies, millas a kilómetros, etc.) para obtener resultados precisos.
  - Criterios de aceptación:
    - La API debe tener un endpoint /longitud/convertir.
    - Debe aceptar unidades de origen y destino, así como el valor a convertir.
    - Debe devolver el resultado de la conversión en formato JSON.
    - Los datos de conversión deben leerse desde un archivo.
2. Como usuario, quiero poder convertir unidades de temperatura (Celsius a Fahrenheit, Kelvin a Celsius, etc.) para entender las mediciones en diferentes escalas.
  - Criterios de aceptación:
    - La API debe tener un endpoint /temperatura/convertir.
    - Debe aceptar unidades de origen y destino, así como el valor a convertir.
    - Debe devolver el resultado de la conversión en formato JSON.
    - Los datos de conversión deben leerse desde un archivo.
3. Como usuario, quiero poder convertir unidades de peso (kilogramos a libras, gramos a onzas, etc.) para aplicaciones de cocina, envíos y más.
  - Criterios de aceptación:
    - La API debe tener un endpoint /peso/convertir.
    - Debe aceptar unidades de origen y destino, así como el valor a convertir.
    - Debe devolver el resultado de la conversión en formato JSON.
    - Los datos de conversión deben leerse desde un archivo.
4. Como usuario, quiero poder convertir unidades de volumen (litros a galones, mililitros a tazas, etc.) para cocinar, medir líquidos y otras tareas.
  - Criterios de aceptación:
    - La API debe tener un endpoint /volumen/convertir.
    - Debe aceptar unidades de origen y destino, así como el valor a convertir.
    - Debe devolver el resultado de la conversión en formato JSON.
    - Los datos de conversión deben leerse desde un archivo.
5. Como desarrollador, quiero que la API maneje errores de manera elegante (unidades no válidas, valores incorrectos) para proporcionar una buena experiencia al usuario.
  - Criterios de aceptación:
    - La API debe devolver códigos de error HTTP apropiados (400, 404, etc.) en caso de errores.
    - Los mensajes de error deben ser claros y descriptivos.
    - La API debe de validar los datos entrantes.
6. Como desarrollador, quiero que los datos de conversión estén almacenados en archivos externos (JSON, CSV) para facilitar la actualización y el mantenimiento.
  - Criterios de aceptación:
    - La API debe leer los datos de conversión desde archivos.
    - Los archivos deben estar organizados por categoría de unidad.
    - La API debe de poder leer multiples formatos de archivos.
7. Como usuario, quiero recibir un mensaje de error claro si proporciono unidades de medida que no existen.
  - Criterios de aceptación:
    - Si el usuario proporciona una unidad de medida que no existe, la API debe devolver un código de error http 400.
    - El mensaje de error debe indicar cual unidad de medida no es valida.
8. Como usuario, quiero recibir un mensaje de error claro si proporciono un valor de medida no valido.
  - Criterios de aceptación:
    - Si el usuario proporciona un valor de medida que no sea numerico, la API debe devolver un codigo de error http 400.
    - El mensaje de error debe de indicar que el valor proporcionado no es valido.

## Consideraciones adicionales:
- Podrías añadir más categorías de unidades (tiempo, velocidad, etc.) según tus necesidades.
- Considera añadir validación de entrada para prevenir errores.
- Implementa pruebas unitarias para cada endpoint y función de conversión.