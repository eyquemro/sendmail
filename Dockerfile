# Étape 1: Construire l'application
FROM golang:1.22 AS builder

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers de modules Go
COPY go.mod go.sum ./

# Télécharger les modules Go
RUN go mod download

# Copier le code source
COPY . .

# Construire l'application
RUN go build -o main .

# Étape 2: Créer l'image finale
FROM gcr.io/distroless/base

# Copier le binaire pré-construit depuis l'étape précédente
COPY --from=builder /app/main /main

# Commande pour exécuter le binaire
ENTRYPOINT ["/main"]
