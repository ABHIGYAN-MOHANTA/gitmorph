

# GitMorph 
**Ahora también disponible en Homebrew**

GitMorph es una potente herramienta de línea de comandos (CLI) que te permite cambiar sin problemas entre múltiples identidades de Git en tu máquina local. Garantiza que todos los comandos de Git utilicen la identidad y la clave SSH del perfil activo. Ideal para desarrolladores que trabajan en diferentes proyectos con varias cuentas de Git y desean mantener consistencia en los commits entre repositorios.

<img width="949" height="982" alt="gitmorph-hero" src="https://github.com/user-attachments/assets/40395ec0-e16a-40b2-8480-25cec1052f31" />

---

## ⚠️ Nota importante: Envoltorio de comandos de Git

A partir de la **v3+**, GitMorph actúa como un **envoltorio para todos los comandos de Git** en repositorios con un perfil de GitMorph activo.

* Ahora, cualquier comando de Git debe tener el prefijo `gitmorph`.
* Ejemplos:

```bash
gitmorph add .
gitmorph commit -m "Your commit message"
gitmorph push origin main
```

* Esto garantiza que la identidad de Git y la clave SSH del perfil activo se apliquen correctamente.
* Ejecutar comandos de `git` sin prefijo puede omitir el perfil activo y provocar commits incorrectos.
* Los usuarios que actualicen deben ejecutar `gitmorph fix` una vez después de la actualización.
* ⚠️ Recomendamos encarecidamente utilizar un alias como `alias gim='gitmorph'`

---

## Características

* Crear y gestionar múltiples perfiles de Git
* Cambiar fácilmente entre diferentes identidades de Git (incluida una clave SSH por perfil)
* Listar todos los perfiles disponibles (muestra la ruta de la clave SSH)
* Editar perfiles existentes
* Eliminar perfiles
* Cambio automático específico por repositorio usando el archivo `.gitmorph`
* Interfaz de línea de comandos simple e intuitiva

---

## Instalación

Asegúrate de tener **Go** instalado y luego ejecuta:

```bash
go install github.com/abhigyan-mohanta/gitmorph@latest
```
o con Homebrew

```bash
brew tap abhigyan-mohanta/homebrew-tap
brew install --cask gitmorph
```

### Actualizar PATH

Después de la instalación, añade los binarios de Go a tu `PATH`:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

---

## Uso

### Crear un nuevo perfil

```bash
gitmorph new
```

Solicita lo siguiente:

* **Nombre del perfil**
* **Nombre de usuario de Git**
* **Correo electrónico de Git**
* **Ruta de la clave privada SSH** (déjalo en blanco para `~/.ssh/id_ed25519`)
* **¿Establecer este perfil como predeterminado?** (y/N)

---

### Listar todos los perfiles

```bash
gitmorph list
```

Salida de ejemplo:

```
Available Git profiles:
NAME      USERNAME          EMAIL                     SSH KEY                     FLAGS
work      abhigyan6602      abhigyan@hostagedown.com  ~/.ssh/id_ed25519           [default]
personal  abhigyan-mohanta  underthunder02@gmail.com  ~/.ssh/id_ed25519_personal  [active]
```

---

### Establecer un perfil predeterminado

```bash
gitmorph default <profile-name>
```

* Establece el perfil especificado como predeterminado global.

Ejemplo:

```bash
gitmorph default personal
# Profile 'personal' is now the default.
```

---

### Activar un perfil específico del repositorio

```bash
gitmorph activate <profile-name>
```

* Establece `user.name` y `user.email` de forma global
* Establece o desactiva `core.sshCommand` para utilizar la clave SSH del perfil
* Añade un archivo `.gitmorph` en la raíz del proyecto para usar este perfil automáticamente

---

### Desactivar un perfil específico del repositorio

```bash
gitmorph deactivate
```

* Desactiva el perfil específico del repositorio
* Vuelve al perfil predeterminado global

Ejemplo:

```bash
gitmorph deactivate
# Project-specific profile deactivated. Falling back to default profile.
# Switched to profile 'work' (global)
```

---

### Reparar tu configuración de GitMorph

```bash
gitmorph fix
```

* Repara `~/.gitmorph.json` si algo salió mal
* Reaplica el perfil predeterminado si es necesario

Ejemplo:

```bash
gitmorph fix
# Fixed ~/.gitmorph.json.
# 'personal' is now the default profile.
# Change it using 'gitmorph default <profile>'
```

*Recomendado para usuarios que migren de la v2 a la v3+.*

---

### Editar un perfil

```bash
gitmorph edit <profile-name>
```

Actualiza de forma interactiva:

* Nombre de usuario
* Correo electrónico
* Ruta de la clave SSH
* Bandera de perfil predeterminado (actual: false)

Deja un campo en blanco para mantener el valor actual.

---

### Eliminar un perfil

```bash
gitmorph delete <profile-name>
```

Elimina la entrada del perfil de `~/.gitmorph.json`.

---

### Nota de migración (v3+)

Si actualizas desde GitMorph v2:

* Ejecuta `gitmorph fix` para migrar tu configuración existente
* Vuelve a establecer tu perfil predeterminado con `gitmorph default <profile>` si es necesario
* Verifica los perfiles específicos del repositorio con `gitmorph deactivate`

---

## Configuración de SSH

Aún puedes usar `~/.ssh/config`; el `core.sshCommand` de GitMorph lo anula cuando está establecido.

Ejemplo:

```plaintext
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519

Host github.com-work
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519_work
```

---

## Dependencias

```go
require (
    github.com/spf13/cobra v1.8.1
    github.com/spf13/pflag v1.0.5
)
```

---

## Contribuciones

¡Las contribuciones son bienvenidas! Por favor, envía un Pull Request.
