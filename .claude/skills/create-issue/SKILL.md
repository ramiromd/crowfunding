---
name: create-issue
description: Crea un nuevo issue en un repositorio de GitHub.
disable-model-invocation: true
---

- Crear un nuevo issue en el repositorio https://github.com/$GITHUB_REPO con el título y el contenido especificados en el archivo $USER_STORIES$ARGUMENTS[0]. El argumento $ARGUMENTS[0] puede ser solo el nombre del archivo o incluir subdirectorios (ej: `user-story.md` o `sprint1/user-story.md`).
- El archivo especificado debe contener el título en la primera línea y el contenido del issue a partir de la segunda línea.
- El archivo especificado debe ser un archivo .md con el formato mencionado anteriormente. En caso de no ser proporcionado o no cumplir con el formato, se debe retornar un error indicando la falta de información o el formato incorrecto.
- El issue debe ser asignado al usuario $ARGUMENTS[1].
- El issue debe ser creado con el milestone $ARGUMENTS[2]. 
- De no existir el milestone, se debe crear uno nuevo con ese nombre y asignarlo al issue. 
- Si el milestone ya existe, se debe asignar al issue sin crear uno nuevo.
- Si el milestone no es especificado, el issue debe ser creado sin asignar ningún milestone.