#define STB_IMAGE_IMPLEMENTATION
#include <glad/glad.h>
#include <GLFW/glfw3.h>
#include <stb_image.h>

#include <glm/glm.hpp>
#include <glm/gtc/matrix_transform.hpp>
#include <glm/gtc/type_ptr.hpp>

#include <shader_m.h>
#include <camera.h>
#include <model.h>

#include <iostream>

// 帧缓冲大小回调函数
void framebuffer_size_callback(GLFWwindow* window, int width, int height);
// 鼠标回调函数
void mouse_callback(GLFWwindow* window, double xpos, double ypos);
// 滚动回调函数
void scroll_callback(GLFWwindow* window, double xoffset, double yoffset);
// 处理输入
void processInput(GLFWwindow* window);
// 加载纹理
unsigned int loadTexture(const char* path);
// 加载立方体贴图
unsigned int loadCubemap(vector<std::string> faces);

// 设置
const unsigned int SCR_WIDTH = 800; // 屏幕宽度
const unsigned int SCR_HEIGHT = 600; // 屏幕高度

// 相机
Camera camera(glm::vec3(0.0f, 0.0f, 3.0f)); // 相机位置
float lastX = (float)SCR_WIDTH / 2.0; // 上一次鼠标x坐标
float lastY = (float)SCR_HEIGHT / 2.0; // 上一次鼠标y坐标
bool firstMouse = true; // 首次鼠标使用

// 计时
float deltaTime = 0.0f; // 每帧时间差
float lastFrame = 0.0f; // 上一帧时间

// lighting
glm::vec3 lightPos(1.2f, 1.0f, 2.0f);

// --- 新增：聚光灯控制变量 (声明为全局变量) ---
float spotlightIntensity = 0.5f;                // 初始亮度/强度因子
glm::vec3 spotlightColor = glm::vec3(1.0f, 1.0f, 1.0f); // 初始颜色 (白色)
int colorIndex = 0;                             // 用于切换预设颜色

// 预设颜色列表
std::vector<glm::vec3> presetSpotlightColors = {
    glm::vec3(1.0f, 1.0f, 1.0f), // 白色
    glm::vec3(1.0f, 0.0f, 0.0f), // 红色
    glm::vec3(0.0f, 1.0f, 0.0f), // 绿色
    glm::vec3(0.0f, 0.0f, 1.0f), // 蓝色
    glm::vec3(1.0f, 1.0f, 0.0f), // 黄色
    glm::vec3(1.0f, 0.5f, 0.0f), // 橙色
    glm::vec3(0.5f, 0.0f, 0.5f)  // 紫色
};

// 用于按键只触发一次的辅助变量
bool K_pressed_last_frame = false; // K键控制颜色
bool L_pressed_last_frame = false; // L键控制重置 (可选)



int main()
{

    // glfw: 初始化和配置
    glfwInit();
    glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3); // 主版本号
    glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 3); // 次版本号
    glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE); // 指定OpenGL核心模式

    // glfw窗口创建
    GLFWwindow* window = glfwCreateWindow(SCR_WIDTH, SCR_HEIGHT, "LearnOpenGL", NULL, NULL);
    if (window == NULL)
    {
        std::cout << "Failed to create GLFW window" << std::endl; // 创建GLFW窗口失败
        glfwTerminate(); // 终止GLFW
        return -1;
    }
    glfwMakeContextCurrent(window); // 将创建的窗口设置为当前的GL上下文
    glfwSetFramebufferSizeCallback(window, framebuffer_size_callback); // 设置帧缓冲大小回调
    glfwSetCursorPosCallback(window, mouse_callback); // 设置鼠标位置回调
    glfwSetScrollCallback(window, scroll_callback); // 设置滚动回调

    // 告诉GLFW捕获鼠标
    glfwSetInputMode(window, GLFW_CURSOR, GLFW_CURSOR_DISABLED);

    // glad: 加载所有OpenGL函数指针
    if (!gladLoadGLLoader((GLADloadproc)glfwGetProcAddress))
    {
        std::cout << "Failed to initialize GLAD" << std::endl; // GLAD初始化失败
        return -1;
    }

    // 配置全局OpenGL状态
    glEnable(GL_DEPTH_TEST); // 启用深度测试

    // 构建并编译着色器
    Shader shader("shader.vs", "shader.fs"); // 立方体贴图着色器
    Shader skyboxShader("skybox.vs", "skybox.fs"); // 天空盒着色器
    Shader lightingShader("multiple_lights.vs", "multiple_lights.fs");
    Shader lightCubeShader("light_cube.vs", "light_cube.fs");





    // 设置顶点数据（和缓冲区）并配置顶点属性
    float cubeVertices[] = {
        // positions          // normals           // texture coords
        -0.5f, -0.5f, -0.5f,  0.0f,  0.0f, -1.0f,  0.0f,  0.0f,
         0.5f, -0.5f, -0.5f,  0.0f,  0.0f, -1.0f,  1.0f,  0.0f,
         0.5f,  0.5f, -0.5f,  0.0f,  0.0f, -1.0f,  1.0f,  1.0f,
         0.5f,  0.5f, -0.5f,  0.0f,  0.0f, -1.0f,  1.0f,  1.0f,
        -0.5f,  0.5f, -0.5f,  0.0f,  0.0f, -1.0f,  0.0f,  1.0f,
        -0.5f, -0.5f, -0.5f,  0.0f,  0.0f, -1.0f,  0.0f,  0.0f,

        -0.5f, -0.5f,  0.5f,  0.0f,  0.0f,  1.0f,  0.0f,  0.0f,
         0.5f, -0.5f,  0.5f,  0.0f,  0.0f,  1.0f,  1.0f,  0.0f,
         0.5f,  0.5f,  0.5f,  0.0f,  0.0f,  1.0f,  1.0f,  1.0f,
         0.5f,  0.5f,  0.5f,  0.0f,  0.0f,  1.0f,  1.0f,  1.0f,
        -0.5f,  0.5f,  0.5f,  0.0f,  0.0f,  1.0f,  0.0f,  1.0f,
        -0.5f, -0.5f,  0.5f,  0.0f,  0.0f,  1.0f,  0.0f,  0.0f,

        -0.5f,  0.5f,  0.5f, -1.0f,  0.0f,  0.0f,  1.0f,  0.0f,
        -0.5f,  0.5f, -0.5f, -1.0f,  0.0f,  0.0f,  1.0f,  1.0f,
        -0.5f, -0.5f, -0.5f, -1.0f,  0.0f,  0.0f,  0.0f,  1.0f,
        -0.5f, -0.5f, -0.5f, -1.0f,  0.0f,  0.0f,  0.0f,  1.0f,
        -0.5f, -0.5f,  0.5f, -1.0f,  0.0f,  0.0f,  0.0f,  0.0f,
        -0.5f,  0.5f,  0.5f, -1.0f,  0.0f,  0.0f,  1.0f,  0.0f,

         0.5f,  0.5f,  0.5f,  1.0f,  0.0f,  0.0f,  1.0f,  0.0f,
         0.5f,  0.5f, -0.5f,  1.0f,  0.0f,  0.0f,  1.0f,  1.0f,
         0.5f, -0.5f, -0.5f,  1.0f,  0.0f,  0.0f,  0.0f,  1.0f,
         0.5f, -0.5f, -0.5f,  1.0f,  0.0f,  0.0f,  0.0f,  1.0f,
         0.5f, -0.5f,  0.5f,  1.0f,  0.0f,  0.0f,  0.0f,  0.0f,
         0.5f,  0.5f,  0.5f,  1.0f,  0.0f,  0.0f,  1.0f,  0.0f,

        -0.5f, -0.5f, -0.5f,  0.0f, -1.0f,  0.0f,  0.0f,  1.0f,
         0.5f, -0.5f, -0.5f,  0.0f, -1.0f,  0.0f,  1.0f,  1.0f,
         0.5f, -0.5f,  0.5f,  0.0f, -1.0f,  0.0f,  1.0f,  0.0f,
         0.5f, -0.5f,  0.5f,  0.0f, -1.0f,  0.0f,  1.0f,  0.0f,
        -0.5f, -0.5f,  0.5f,  0.0f, -1.0f,  0.0f,  0.0f,  0.0f,
        -0.5f, -0.5f, -0.5f,  0.0f, -1.0f,  0.0f,  0.0f,  1.0f,

        -0.5f,  0.5f, -0.5f,  0.0f,  1.0f,  0.0f,  0.0f,  1.0f,
         0.5f,  0.5f, -0.5f,  0.0f,  1.0f,  0.0f,  1.0f,  1.0f,
         0.5f,  0.5f,  0.5f,  0.0f,  1.0f,  0.0f,  1.0f,  0.0f,
         0.5f,  0.5f,  0.5f,  0.0f,  1.0f,  0.0f,  1.0f,  0.0f,
        -0.5f,  0.5f,  0.5f,  0.0f,  1.0f,  0.0f,  0.0f,  0.0f,
        -0.5f,  0.5f, -0.5f,  0.0f,  1.0f,  0.0f,  0.0f,  1.0f
    };
    float skyboxVertices[] = {
        // 位置
        -1.0f,  1.0f, -1.0f,
        -1.0f, -1.0f, -1.0f,
         1.0f, -1.0f, -1.0f,
         1.0f, -1.0f, -1.0f,
         1.0f,  1.0f, -1.0f,
        -1.0f,  1.0f, -1.0f,

        -1.0f, -1.0f,  1.0f,
        -1.0f, -1.0f, -1.0f,
        -1.0f,  1.0f, -1.0f,
        -1.0f,  1.0f, -1.0f,
        -1.0f,  1.0f,  1.0f,
        -1.0f, -1.0f,  1.0f,

         1.0f, -1.0f, -1.0f,
         1.0f, -1.0f,  1.0f,
         1.0f,  1.0f,  1.0f,
         1.0f,  1.0f,  1.0f,
         1.0f,  1.0f, -1.0f,
         1.0f, -1.0f, -1.0f,

        -1.0f, -1.0f,  1.0f,
        -1.0f,  1.0f,  1.0f,
         1.0f,  1.0f,  1.0f,
         1.0f,  1.0f,  1.0f,
         1.0f, -1.0f,  1.0f,
        -1.0f, -1.0f,  1.0f,

        -1.0f,  1.0f, -1.0f,
         1.0f,  1.0f, -1.0f,
         1.0f,  1.0f,  1.0f,
         1.0f,  1.0f,  1.0f,
        -1.0f,  1.0f,  1.0f,
        -1.0f,  1.0f, -1.0f,

        -1.0f, -1.0f, -1.0f,
        -1.0f, -1.0f,  1.0f,
         1.0f, -1.0f, -1.0f,
         1.0f, -1.0f, -1.0f,
        -1.0f, -1.0f,  1.0f,
         1.0f, -1.0f,  1.0f
    };


    // positions all containers
    glm::vec3 cubePositions[] = {
        glm::vec3(2.0f,  2.0f,  2.0f),
        glm::vec3(2.0f,  5.0f, -15.0f),
        glm::vec3(-1.5f, -2.2f, -2.5f),
        glm::vec3(-3.8f, -2.0f, -12.3f),
        glm::vec3(2.4f, -0.4f, -3.5f),
        glm::vec3(-1.7f,  3.0f, -7.5f),
        glm::vec3(1.3f, -2.0f, -2.5f),
        glm::vec3(1.5f,  2.0f, -2.5f),
        glm::vec3(1.5f,  0.2f, -1.5f),
        glm::vec3(-1.3f,  1.0f, -1.5f)
    };


    // positions of the point lights
    glm::vec3 pointLightPositions[] = {
        glm::vec3(0.7f,  0.2f,  2.0f),
        glm::vec3(2.3f, -3.3f, -4.0f),
        glm::vec3(-4.0f,  2.0f, -12.0f),
        glm::vec3(0.0f,  0.0f, -3.0f)
    };

    


    // 创建立方体VAO
    unsigned int VBO, cubeVAO;
    glGenVertexArrays(1, &cubeVAO);
    glGenBuffers(1, &VBO);

    glBindBuffer(GL_ARRAY_BUFFER, VBO);
    glBufferData(GL_ARRAY_BUFFER, sizeof(cubeVertices), cubeVertices, GL_STATIC_DRAW);

    glBindVertexArray(cubeVAO);
    glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 8 * sizeof(float), (void*)0);
    glEnableVertexAttribArray(0);
    glVertexAttribPointer(1, 3, GL_FLOAT, GL_FALSE, 8 * sizeof(float), (void*)(3 * sizeof(float)));
    glEnableVertexAttribArray(1);
    glVertexAttribPointer(2, 2, GL_FLOAT, GL_FALSE, 8 * sizeof(float), (void*)(6 * sizeof(float)));
    glEnableVertexAttribArray(2);

    // 创建天空盒VAO
    unsigned int skyboxVAO, skyboxVBO;
    glGenVertexArrays(1, &skyboxVAO);
    glGenBuffers(1, &skyboxVBO);
    glBindVertexArray(skyboxVAO);
    glBindBuffer(GL_ARRAY_BUFFER, skyboxVBO);
    glBufferData(GL_ARRAY_BUFFER, sizeof(skyboxVertices), &skyboxVertices, GL_STATIC_DRAW);
    glEnableVertexAttribArray(0);
    glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);

    // 配置光源的VAO（VBO保持不变；顶点与光源对象相同，也是一个3D立方体）
    unsigned int lightCubeVAO;
    glGenVertexArrays(1, &lightCubeVAO);
    glBindVertexArray(lightCubeVAO);

    glBindBuffer(GL_ARRAY_BUFFER, VBO);
    // note that we update the lamp's position attribute's stride to reflect the updated buffer data
    glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 8 * sizeof(float), (void*)0);
    glEnableVertexAttribArray(0);

    // load textures (we now use a utility function to keep the code more organized)
    // -----------------------------------------------------------------------------
    unsigned int diffuseMap = loadTexture("texture2.jpg");
    unsigned int specularMap = loadTexture("texture3.png");

    // shader configuration
    // --------------------
    lightingShader.use();
    lightingShader.setInt("material.diffuse", 1);
    lightingShader.setInt("material.specular", 1);


    // 加载纹理
    vector<std::string> faces{
        "posx.jpg", // 右
        "negx.jpg", // 左
        "posy.jpg", // 上
        "negy.jpg", // 下
        "posz.jpg", // 前
        "negz.jpg", // 后
    };
    unsigned int cubemapTexture = loadCubemap(faces);

    // 着色器配置
    // --------------------
    lightingShader.use();
    lightingShader.setInt("material.diffuse", 0);
    lightingShader.setInt("material.specular", 1);

    // 着色器配置
    shader.use(); // 使用立方体贴图着色器
    shader.setInt("skybox", 0); // 设置天空盒纹理为纹理单元0

    skyboxShader.use(); // 使用天空盒着色器
    skyboxShader.setInt("skybox", 0); // 设置天空盒纹理为纹理单元0
    // view/projection transformations
    glm::mat4 projection;
    glm::mat4 view;

    while (!glfwWindowShouldClose(window))
    {
        // 每帧时间逻辑
        float currentFrame = static_cast<float>(glfwGetTime());
        deltaTime = currentFrame - lastFrame;
        lastFrame = currentFrame;

        // 输入处理
        processInput(window);

        // 渲染
        glClearColor(0.1f, 0.1f, 0.1f, 1.0f); // 设置清除颜色
        glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT); // 清除颜色和深度缓冲区

        // 正常绘制场景
        shader.use();
        glm::mat4 model = glm::mat4(1.0f); // 模型矩阵
        shader.setMat4("model", model); // 设置着色器模型矩阵
        shader.setMat4("view", view); // 设置着色器视图矩阵
        shader.setMat4("projection", projection); // 设置着色器投影矩阵
        shader.setVec3("cameraPos", camera.Position); // 设置着色器相机位置

        // 绘制立方体
        glBindVertexArray(cubeVAO);
        glActiveTexture(GL_TEXTURE0);
        glBindTexture(GL_TEXTURE_CUBE_MAP, cubemapTexture);
        glDrawArrays(GL_TRIANGLES, 0, 36);
        glBindVertexArray(0);

        // 最后绘制天空盒
        glDepthFunc(GL_LEQUAL); // 修改深度函数
        skyboxShader.use();
        view = glm::mat4(glm::mat3(camera.GetViewMatrix())); // 从视图矩阵中移除平移
        skyboxShader.setMat4("view", view); // 设置天空盒着色器视图矩阵
        skyboxShader.setMat4("projection", projection); // 设置天空盒着色器投影矩阵

        // 绘制天空盒立方体
        glBindVertexArray(skyboxVAO);
        glActiveTexture(GL_TEXTURE0);
        glBindTexture(GL_TEXTURE_CUBE_MAP, cubemapTexture);
        glDrawArrays(GL_TRIANGLES, 0, 36);
        glBindVertexArray(0);
        glDepthFunc(GL_LESS); // 恢复默认深度函数

        // activate shader
        lightingShader.use();
        lightingShader.setVec3("viewPos", camera.Position);
        lightingShader.setFloat("material.shininess", 36.0f);

      
        // directional light
        lightingShader.setVec3("dirLight.direction", -0.2f, -1.0f, -0.3f);
        lightingShader.setVec3("dirLight.ambient", 0.1f, 0.1f, 0.1f); // 增加环境光强度
        lightingShader.setVec3("dirLight.diffuse", 0.8f, 0.8f, 0.8f); // 增加漫反射强度
        lightingShader.setVec3("dirLight.specular", 1.0f, 1.0f, 1.0f); // 增加镜面反射强度

        // point light 1
        lightingShader.setVec3("pointLights[0].position", pointLightPositions[0]);
        lightingShader.setVec3("pointLights[0].ambient", 0.15f, 0.15f, 0.5f);// 增加环境光强度
        lightingShader.setVec3("pointLights[0].diffuse", 0.8f, 0.8f, 0.8f);// 增加漫反射强度
        lightingShader.setVec3("pointLights[0].specular", 1.0f, 1.0f, 1.0f); // 增加镜面反射强度
        lightingShader.setFloat("pointLights[0].constant", 1.0f);
        lightingShader.setFloat("pointLights[0].linear", 0.09f);
        lightingShader.setFloat("pointLights[0].quadratic", 0.032f);
        // point light 2
        lightingShader.setVec3("pointLights[1].position", pointLightPositions[1]);
        lightingShader.setVec3("pointLights[1].ambient", 0.05f, 0.05f, 0.05f);
        lightingShader.setVec3("pointLights[1].diffuse", 0.8f, 0.8f, 0.8f);
        lightingShader.setVec3("pointLights[1].specular", 1.0f, 1.0f, 1.0f);
        lightingShader.setFloat("pointLights[1].constant", 1.0f);
        lightingShader.setFloat("pointLights[1].linear", 0.09f);
        lightingShader.setFloat("pointLights[1].quadratic", 0.032f);
        // point light 3
        lightingShader.setVec3("pointLights[2].position", pointLightPositions[2]);
        lightingShader.setVec3("pointLights[2].ambient", 0.05f, 0.05f, 0.05f);
        lightingShader.setVec3("pointLights[2].diffuse", 0.8f, 0.8f, 0.8f);
        lightingShader.setVec3("pointLights[2].specular", 1.0f, 1.0f, 1.0f);
        lightingShader.setFloat("pointLights[2].constant", 1.0f);
        lightingShader.setFloat("pointLights[2].linear", 0.09f);
        lightingShader.setFloat("pointLights[2].quadratic", 0.032f);
        // point light 4
        lightingShader.setVec3("pointLights[3].position", pointLightPositions[3]);
        lightingShader.setVec3("pointLights[3].ambient", 0.05f, 0.05f, 0.05f);
        lightingShader.setVec3("pointLights[3].diffuse", 0.8f, 0.8f, 0.8f);
        lightingShader.setVec3("pointLights[3].specular", 1.0f, 1.0f, 1.0f);
        lightingShader.setFloat("pointLights[3].constant", 1.0f);
        lightingShader.setFloat("pointLights[3].linear", 0.09f);
        lightingShader.setFloat("pointLights[3].quadratic", 0.032f);
        // spotLight
        lightingShader.use(); // 确保在使用 lightingShader 前调用
        lightingShader.setVec3("spotLight.position", camera.Position);
        lightingShader.setVec3("spotLight.direction", camera.Front);

        // 定义聚光灯的基础光照颜色/强度 (这些可以根据喜好调整)
        glm::vec3 baseAmbientSpot = glm::vec3(0.0f, 0.0f, 0.0f); // 聚光灯通常环境光较弱或为0
        glm::vec3 baseDiffuseSpot = glm::vec3(0.8f, 0.8f, 0.8f); // 基础漫反射强度
        glm::vec3 baseSpecularSpot = glm::vec3(1.0f, 1.0f, 1.0f); // 基础镜面反射强度

        // 使用全局变量 spotlightColor 和 spotlightIntensity 来计算最终的光照值
        lightingShader.setVec3("spotLight.ambient", baseAmbientSpot * spotlightColor * spotlightIntensity);
        lightingShader.setVec3("spotLight.diffuse", baseDiffuseSpot * spotlightColor * spotlightIntensity);
        lightingShader.setVec3("spotLight.specular", baseSpecularSpot * spotlightColor * spotlightIntensity);

        lightingShader.setFloat("spotLight.constant", 1.0f);
        lightingShader.setFloat("spotLight.linear", 0.09f);
        lightingShader.setFloat("spotLight.quadratic", 0.032f);
        lightingShader.setFloat("spotLight.cutOff", glm::cos(glm::radians(12.5f)));
        lightingShader.setFloat("spotLight.outerCutOff", glm::cos(glm::radians(15.0f)));

        // world transformation
        lightingShader.setMat4("model", model);

        // bind diffuse map
        glActiveTexture(GL_TEXTURE0);
        glBindTexture(GL_TEXTURE_2D, diffuseMap);
        // bind specular map
        glActiveTexture(GL_TEXTURE1);
        glBindTexture(GL_TEXTURE_2D, specularMap);

        // update view/projection transformations
        projection = glm::perspective(glm::radians(camera.Zoom), (float)SCR_WIDTH / (float)SCR_HEIGHT, 0.1f, 100.0f);
        view = camera.GetViewMatrix();
        lightingShader.setMat4("projection", projection);
        lightingShader.setMat4("view", view);

        // render containers
        glBindVertexArray(cubeVAO);
        for (unsigned int i = 0; i < 10; i++)
        {
            glm::mat4 model = glm::mat4(1.0f);
            model = glm::translate(model, cubePositions[i]);
            float angle = 20.0f * i;
            model = glm::rotate(model, glm::radians(angle), glm::vec3(1.0f, 0.3f, 0.5f));
            lightingShader.setMat4("model", model);

            glDrawArrays(GL_TRIANGLES, 0, 36);
        }

        // render point lights
        lightCubeShader.use();
        lightCubeShader.setMat4("projection", projection);
        lightCubeShader.setMat4("view", view);

        glBindVertexArray(lightCubeVAO);
        for (unsigned int i = 0; i < 4; i++)
        {
            glm::mat4 model = glm::mat4(1.0f);
            model = glm::translate(model, pointLightPositions[i]);
            model = glm::scale(model, glm::vec3(0.1f));
            lightCubeShader.setMat4("model", model);

            glDrawArrays(GL_TRIANGLES, 0, 36);
        }


        

        // glfw: 交换缓冲区并轮询IO事件
        glfwSwapBuffers(window);
        glfwPollEvents();
    }



    glfwTerminate(); // 终止GLFW
    return 0;
}

/* 处理所有输入：查询GLFW当前帧是否按下 / 释放了相关键，并相应地作出反应
void processInput(GLFWwindow* window)
{
    if (glfwGetKey(window, GLFW_KEY_ESCAPE) == GLFW_PRESS)
        glfwSetWindowShouldClose(window, true);

    if (glfwGetKey(window, GLFW_KEY_W) == GLFW_PRESS)
        camera.ProcessKeyboard(FORWARD, deltaTime);
    if (glfwGetKey(window, GLFW_KEY_S) == GLFW_PRESS)
        camera.ProcessKeyboard(BACKWARD, deltaTime);
    if (glfwGetKey(window, GLFW_KEY_A) == GLFW_PRESS)
        camera.ProcessKeyboard(LEFT, deltaTime);
    if (glfwGetKey(window, GLFW_KEY_D) == GLFW_PRESS)
        camera.ProcessKeyboard(RIGHT, deltaTime);

}*/

// processInput 函数 (包含聚光灯控制)
void processInput(GLFWwindow* window)
{
    if (glfwGetKey(window, GLFW_KEY_ESCAPE) == GLFW_PRESS)
        glfwSetWindowShouldClose(window, true);

    // 相机控制
    if (glfwGetKey(window, GLFW_KEY_W) == GLFW_PRESS)
        camera.ProcessKeyboard(FORWARD, deltaTime);
    if (glfwGetKey(window, GLFW_KEY_S) == GLFW_PRESS)
        camera.ProcessKeyboard(BACKWARD, deltaTime);
    if (glfwGetKey(window, GLFW_KEY_A) == GLFW_PRESS)
        camera.ProcessKeyboard(LEFT, deltaTime);
    if (glfwGetKey(window, GLFW_KEY_D) == GLFW_PRESS)
        camera.ProcessKeyboard(RIGHT, deltaTime);

    // --- 聚光灯控制 ---
    // 亮度控制
    float intensityStep = 2.5f * deltaTime; // 亮度变化速率，可以调整
    if (glfwGetKey(window, GLFW_KEY_UP) == GLFW_PRESS)
    {
        spotlightIntensity += intensityStep;
        if (spotlightIntensity > 10.0f) spotlightIntensity = 10.0f; // 限制最大亮度
    }
    if (glfwGetKey(window, GLFW_KEY_DOWN) == GLFW_PRESS)
    {
        spotlightIntensity -= intensityStep;
        if (spotlightIntensity < 0.0f) spotlightIntensity = 0.0f; // 限制最小亮度
    }

    // 颜色控制 (按一次K键切换一次)
    bool K_currently_pressed = glfwGetKey(window, GLFW_KEY_K) == GLFW_PRESS;
    if (K_currently_pressed && !K_pressed_last_frame)
    {
        colorIndex = (colorIndex + 1) % presetSpotlightColors.size();
        spotlightColor = presetSpotlightColors[colorIndex];
        std::cout << "Spotlight Color Changed To: Index " << colorIndex
            << " (R:" << spotlightColor.r
            << " G:" << spotlightColor.g
            << " B:" << spotlightColor.b << ")" << std::endl;
    }
    K_pressed_last_frame = K_currently_pressed;

    // (可选) L键重置
    bool L_currently_pressed = glfwGetKey(window, GLFW_KEY_L) == GLFW_PRESS;
    if (L_currently_pressed && !L_pressed_last_frame)
    {
        spotlightIntensity = 1.0f;
        colorIndex = 0;
        spotlightColor = presetSpotlightColors[colorIndex];
        std::cout << "Spotlight Reset!" << std::endl;
    }
    L_pressed_last_frame = L_currently_pressed;
}

// 当窗口大小改变时触发的回调函数
void framebuffer_size_callback(GLFWwindow* window, int width, int height)
{
    // 确保视口与新的窗口尺寸匹配
    glViewport(0, 0, width, height);
}

// 当鼠标移动时触发的回调函数
void mouse_callback(GLFWwindow* window, double xposIn, double yposIn)
{
    float xpos = static_cast<float>(xposIn);
    float ypos = static_cast<float>(yposIn);
    if (firstMouse)
    {
        lastX = xpos;
        lastY = ypos;
        firstMouse = false;
    }

    float xoffset = xpos - lastX; // 鼠标水平偏移量
    float yoffset = lastY - ypos; // 鼠标垂直偏移量，取反因为y坐标从底部到顶部

    lastX = xpos;
    lastY = ypos;

    camera.ProcessMouseMovement(xoffset, yoffset); // 处理鼠标移动
}

// 当鼠标滚轮滚动时触发的回调函数
void scroll_callback(GLFWwindow* window, double xoffset, double yoffset)
{
    camera.ProcessMouseScroll(static_cast<float>(yoffset)); // 处理鼠标滚轮滚动
}

// 从文件加载2D纹理的工具函数
unsigned int loadTexture(char const* path)
{
    unsigned int textureID;
    glGenTextures(1, &textureID); // 创建纹理ID

    int width, height, nrComponents;
    unsigned char* data = stbi_load(path, &width, &height, &nrComponents, 0); // 加载图像数据
    if (data)
    {
        GLenum format;
        if (nrComponents == 1)
            format = GL_RED;
        else if (nrComponents == 3)
            format = GL_RGB;
        else if (nrComponents == 4)
            format = GL_RGBA;

        glBindTexture(GL_TEXTURE_2D, textureID); // 绑定纹理
        glTexImage2D(GL_TEXTURE_2D, 0, format, width, height, 0, format, GL_UNSIGNED_BYTE, data); // 指定2D纹理图像
        glGenerateMipmap(GL_TEXTURE_2D); // 生成纹理的各层mipmap

        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_REPEAT); // 设置纹理水平环绕方式
        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_REPEAT); // 设置纹理垂直环绕方式
        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR_MIPMAP_LINEAR); // 设置纹理的最小过滤方式
        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR); // 设置纹理的放大过滤方式

        stbi_image_free(data); // 释放图像数据
    }
    else
    {
        std::cout << "Texture failed to load at path: " << path << std::endl; // 纹理加载失败
        stbi_image_free(data);
    }

    return textureID; // 返回纹理ID
}

// 从6个独立纹理面加载立方体贴图
unsigned int loadCubemap(vector<std::string> faces)
{
    unsigned int textureID;
    glGenTextures(1, &textureID); // 创建纹理ID
    glBindTexture(GL_TEXTURE_CUBE_MAP, textureID); // 绑定立方体贴图

    int width, height, nrComponents;
    for (unsigned int i = 0; i < faces.size(); i++)
    {
        unsigned char* data = stbi_load(faces[i].c_str(), &width, &height, &nrComponents, 0); // 加载图像数据
        if (data)
        {
            glTexImage2D(GL_TEXTURE_CUBE_MAP_POSITIVE_X + i, 0, GL_RGB, width, height, 0, GL_RGB, GL_UNSIGNED_BYTE, data); // 指定立方体贴图图像
            stbi_image_free(data); // 释放图像数据
        }
        else
        {
            std::cout << "Cubemap texture failed to load at path: " << faces[i] << std::endl; // 立方体贴图加载失败
            stbi_image_free(data);
        }
    }
    glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_MIN_FILTER, GL_LINEAR); // 设置立方体贴图的最小过滤方式
    glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_MAG_FILTER, GL_LINEAR); // 设置立方体贴图的放大过滤方式
    glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_WRAP_S, GL_CLAMP_TO_EDGE); // 设置立方体贴图的水平环绕方式
    glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_WRAP_T, GL_CLAMP_TO_EDGE); // 设置立方体贴图的垂直环绕方式
    glTexParameteri(GL_TEXTURE_CUBE_MAP, GL_TEXTURE_WRAP_R, GL_CLAMP_TO_EDGE); // 设置立方体贴图的深度环绕方式

    return textureID; // 返回纹理ID
}


