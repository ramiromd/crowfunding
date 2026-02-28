---
name: create-issue
description: Crea un nuevo issue en un repositorio de GitHub.
disable-model-invocation: true
---

- Crear un nuevo issue en el repositorio https://github.com/$GITHUB_REPO con el título y el contenido especificados en el archivo $USER_STORIES$ARGUMENTS[0]. El argumento $ARGUMENTS[0] puede ser solo el nombre del archivo o incluir subdirectorios (ej: `user-story.md` o `sprint1/user-story.md`).
- El archivo especificado debe contener el título en la primera línea y el contenido del issue a partir de la segunda línea.
- El archivo especificado debe ser un archivo .md con el formato mencionado anteriormente. En caso de no ser proporcionado o no cumplir con el formato, se debe retornar un error indicando la falta de información o el formato incorrecto.
- El issue debe ser asignado al usuario $ARGUMENTS[1].
- El issue debe ser creado con el milestone identificado por el campo `milestone_id` del front matter YAML del archivo. Este campo contiene el ID numérico del milestone en GitHub (ej: `milestone_id: 1`). El front matter se encuentra al inicio del archivo, delimitado por `---`.
- Es requerido que el milestone exista previamente en el repositorio. Si el milestone no existe, se debe retornar un error indicando que el milestone no fue encontrado.
- Si el campo `milestone_id` no está presente en el front matter, o el archivo no tiene front matter, el issue debe ser creado sin asignar ningún milestone.