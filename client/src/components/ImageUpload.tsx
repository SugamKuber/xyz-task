// ImageUpload.tsx
import React, { useState, ChangeEvent } from 'react';
import { useForm } from 'react-hook-form';

interface UploadResult {
  type: string;
  color: string;
  make: string;
  model: string;
  caption: string;
}

const ImageUpload: React.FC = () => {
  const [image, setImage] = useState<File | null>(null);
  const [result, setResult] = useState<UploadResult | null>(null);
  const { register, handleSubmit, reset } = useForm();

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      setImage(e.target.files[0]);
    }
  };

  const onSubmit = async () => {
    if (!image) return;
    
    const formData = new FormData();
    formData.append('image', image);

    try {
      const response = await fetch('http://localhost:8080/api/upload', {
        method: 'POST',
        body: formData,
      });
      
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      const data: UploadResult = await response.json();
      setResult(data);
    } catch (error) {
      console.error('Error uploading image:', error);
    } finally {
      reset();  
      setImage(null); 
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-4 bg-white rounded-lg shadow-md">
      <form onSubmit={(e) => {
        e.preventDefault();
        handleSubmit(onSubmit)();
      }}>
        <div className="mb-4">
          <input
            type="file"
            accept="image/*"
            onChange={handleFileChange}
            className="block w-full text-sm text-gray-500 file:py-2 file:px-4 file:border file:border-gray-300 file:rounded file:text-sm file:font-medium file:bg-gray-50 file:text-gray-700 hover:file:bg-gray-100"
          />
        </div>
        <button
          type="submit"
          className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600"
        >
          Upload
        </button>
      </form>
      {result && (
        <div className="mt-6 p-4 border border-gray-300 rounded-md">
          <h2 className="text-xl font-semibold mb-2">Upload Result</h2>
          <p><strong>Type:</strong> {result.type}</p>
          <p><strong>Color:</strong> {result.color}</p>
          <p><strong>Make:</strong> {result.make}</p>
          <p><strong>Model:</strong> {result.model}</p>
          <p><strong>Caption:</strong> {result.caption}</p>
        </div>
      )}
    </div>
  );
};

export default ImageUpload;
