def label = "slave-${UUID.randomUUID().toString()}"

podTemplate(label: label, containers: [
  containerTemplate(name: 'golang', image: '172.16.140.21/heyang/golang:1.14.2', command: 'cat', ttyEnabled: true),
  containerTemplate(name: 'docker', image: '172.16.140.21/heyang/docker-cli:19.03.8', command: 'cat', ttyEnabled: true),
  containerTemplate(name: 'kubectl', image: '172.16.140.21/heyang/kubectl:v1.15.3', command: 'cat', ttyEnabled: true),
  containerTemplate(name: 'jnlp', image: '172.16.140.21/heyang/jnlp-slave:4.0.1-1', alwaysPullImage: false, privileged: true, args: '${computer.jnlpmac} ${computer.name}')
], serviceAccount: 'jenkins', volumes: [
  hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock')
]) {
  node(label) {
    def myRepo = checkout scm
    def gitCommit = myRepo.GIT_COMMIT
    def gitBranch = myRepo.GIT_BRANCH
    def imageTag = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()
    def imageUri = "172.16.140.21"
    def imageHub = "heyang"
    def imageApp = "wisdom-client"
    def image = "${imageUri}/${imageHub}/${imageApp}"
    def site = ""

    stage('单元测试') {
      echo "Part1.单元测试-test"
    }
    stage('代码编译打包') {
      try {
        container('golang') {
            echo "Part2.代码编译打包"
            sh "pwd"
            sh "ls -l"
            sh '''
                export GO111MODULE=off
                go env
                cp -Ra ../wisdomClient/ /go/src/
                cp -Ra ../wisdomClient/vendor/* /go/src/
                mv /go/src/wisdomClient/ /go/src/wisdom-client
                ls /go/src
                make build
            '''
        }
      } catch (exc) {
        println "构建失败 - ${currentBuild.fullDisplayName}"
        throw(exc)
      }
    }
    stage('构建Docker镜像') {
      withCredentials([usernamePassword(credentialsId: 'heyang-harbor-auth', passwordVariable: 'DOCKER_HUB_PASSWORD', usernameVariable: 'DOCKER_HUB_USER')]) {
        container('docker') {
          echo "Part3.构建Docker镜像"
          def userInput = input(
            id: 'userInput',
            message: '选择一个部署环境',
            parameters: [
              [
                $class: 'ChoiceParameterDefinition',
                choices: "TJ\nNJ\nHZ",
                name: 'EnvConfig'
              ]
            ]
          )
          echo "部署应用到 ${userInput} 配置文件"
          if (userInput == "TJ") {
            sh """
                docker login ${imageUri} -u ${DOCKER_HUB_USER} -p ${DOCKER_HUB_PASSWORD}
                docker build --build-arg CONFIG=./conf/tj-config.yaml -t ${image}:${imageTag} .
                docker push ${image}:${imageTag}
            """
          } else if (userInput == "NJ"){
            sh """
                docker login ${imageUri} -u ${DOCKER_HUB_USER} -p ${DOCKER_HUB_PASSWORD}
                docker build --build-arg CONFIG=./conf/nj-config.yaml -t ${image}:${imageTag} .
                docker push ${image}:${imageTag}
            """
          } else {
            sh """
                docker login ${imageUri} -u ${DOCKER_HUB_USER} -p ${DOCKER_HUB_PASSWORD}
                docker build --build-arg CONFIG=./conf/hz-config.yaml -t ${image}:${imageTag} .
                docker push ${image}:${imageTag}
            """
          }
          site = userInput
        }
      }
    }
    stage('修改部署文件') {
      echo "Part4.修改YAML文件参数"
      def ciEnv = "dev"
      if (gitBranch == "origin/master") {
        ciEnv = "prod"
      }
      sh "echo ${site}"
      sh "sed -i 's/<IMAGE_URI>/${imageUri}/g' manifests/deployment.yaml"
      sh "sed -i 's/<IMAGE_HUB>/${imageHub}/g' manifests/deployment.yaml"
      sh "sed -i 's/<IMAGE_APP>/${imageApp}/g' manifests/deployment.yaml"
      sh "sed -i 's/<BUILD_TAG>/${imageTag}/g' manifests/deployment.yaml"
      sh "sed -i 's/<BRANCH_NAME>/${ciEnv}/g' manifests/deployment.yaml"
    }
    stage('推送Kubernetes') {
      withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
        container('kubectl') {
         sh "mkdir -p ~/.kube && cp ${KUBECONFIG} ~/.kube/config"
         echo "Part5.部署应用到 K8S"
         sh '''
             kubectl config use-context tj-k8s
             kubectl apply -f manifests/deployment.yaml
             kubectl apply -f manifests/service.yaml
             kubectl apply -f manifests/ingress.yaml
             kubectl rollout status -f manifests/deployment.yaml
         '''
         echo "6.部署成功"
        }
      }
    }
  }
}