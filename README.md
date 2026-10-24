# Hello Go в Minikube

Приложение разворачивается в Minikube с помощью Helm. Внешний HTTP-трафик принимает NGINX Ingress Controller, а само приложение доступно внутри кластера через Service типа `ClusterIP`.

## Требования

- Docker Desktop запущен;
- установлен `minikube` и создан кластер;
- установлены `kubectl` и `helm`;
- включён Ingress addon:

  ```powershell
  minikube addons enable ingress
  ```

## Запуск

1. Соберите Docker-образ приложения:

   ```powershell
   docker build -t hello-go:1.3 .
   ```

2. Загрузите образ в Minikube:

   ```powershell
   minikube image load hello-go:1.3
   ```

   Тег образа должен совпадать со значением `image.tag` в `httpserver/values.yaml`. Если тег изменён, пересоберите и загрузите образ с новым тегом.

3. Проверьте Helm chart до установки:

   ```powershell
   helm template hello-go .\httpserver
   helm lint .\httpserver
   ```

4. Установите или обновите Helm-релиз:

   ```powershell
   helm upgrade --install hello-go .\httpserver
   ```

5. В отдельном окне запустите туннель Minikube:

   ```powershell
   minikube tunnel
   ```

   Туннель должен оставаться запущенным, пока сервис нужен из браузера. После его запуска приложение доступно по адресу [http://localhost](http://localhost).

## Проверка

Проверьте Helm-релиз:

```powershell
helm list
helm status hello-go
```

Проверьте Kubernetes-объекты:

```powershell
kubectl get pods,deployments,services,configmaps,secrets,ingress
```

Pod должен иметь статус `Running`, а Ingress должен содержать адрес `localhost`.
